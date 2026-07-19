import requests
import json
import sys


BASE_URL = "http://localhost:8080/api/v1"



def get_field(data, field):
    """
    Supports Go default JSON names and snake_case names.

    Examples:
        id            -> id / ID
        project_id    -> project_id / ProjectID
        story_points  -> story_points / StoryPoints
        sprint_id     -> sprint_id / SprintID
    """

    # Exact match
    if field in data:
        return data[field]


    # Handle common Go initialisms
    parts = field.split("_")

    pascal_parts = []

    for part in parts:

        if part.lower() == "id":
            pascal_parts.append("ID")

        elif part.lower() == "url":
            pascal_parts.append("URL")

        elif part.lower() == "api":
            pascal_parts.append("API")

        elif part.lower() == "http":
            pascal_parts.append("HTTP")

        else:
            pascal_parts.append(
                part.capitalize()
            )


    pascal = "".join(pascal_parts)


    if pascal in data:
        return data[pascal]


    raise KeyError(
        f"Missing field '{field}' in response:\n{data}"
    )


def check(response, name):

    print("\n====================")
    print(name)
    print("====================")

    print("STATUS:", response.status_code)

    try:
        print(
            json.dumps(
                response.json(),
                indent=2
            )
        )

    except Exception:
        print(response.text)


    if response.status_code >= 400:

        print("\nFAILED")
        sys.exit(1)


    print("PASSED")



def assert_field(data, field, expected, name):

    actual = get_field(data, field)

    if actual != expected:

        print("\nFAILED:", name)
        print("Expected:", expected)
        print("Actual:", actual)

        sys.exit(1)


    print("Verified:", name)



# =====================================================
# CREATE PROJECT
# =====================================================

project = requests.post(
    f"{BASE_URL}/projects",
    json={
        "name": "Project Hive Test",
        "key": "PHT",
        "description": "Integration test project"
    }
)

check(project, "Create Project")


project_data = project.json()


project_id = get_field(
    project_data,
    "id"
)


assert_field(
    project_data,
    "name",
    "Project Hive Test",
    "Project name"
)


assert_field(
    project_data,
    "key",
    "PHT",
    "Project key"
)



# =====================================================
# CREATE EPIC
# =====================================================

epic = requests.post(
    f"{BASE_URL}/projects/{project_id}/epics",
    json={
        "name": "Authentication",
        "description": "Authentication features"
    }
)

check(epic, "Create Epic")


epic_data = epic.json()


epic_id = get_field(
    epic_data,
    "id"
)


assert_field(
    epic_data,
    "name",
    "Authentication",
    "Epic name"
)


assert_field(
    epic_data,
    "project_id",
    project_id,
    "Epic project relationship"
)



# =====================================================
# CREATE ISSUE WITHOUT EPIC
# =====================================================

issue = requests.post(
    f"{BASE_URL}/projects/{project_id}/issues",
    json={
        "title": "Build Login Screen",
        "description": "Create login UI",
        "type": "story",
        "priority": "high",
        "story_points": 5
    }
)

check(issue, "Create Backlog Issue")


issue_data = issue.json()


issue_id = get_field(
    issue_data,
    "id"
)


assert_field(
    issue_data,
    "title",
    "Build Login Screen",
    "Issue title"
)


assert_field(
    issue_data,
    "project_id",
    project_id,
    "Issue project relationship"
)


assert_field(
    issue_data,
    "story_points",
    5,
    "Issue story points"
)



# =====================================================
# CREATE ISSUE WITH EPIC
# =====================================================

issue2 = requests.post(
    f"{BASE_URL}/projects/{project_id}/issues",
    json={
        "title": "Password Reset",
        "description": "Reset password workflow",
        "type": "task",
        "priority": "medium",
        "story_points": 3,
        "epic_id": epic_id
    }
)

check(issue2, "Create Epic Issue")


issue2_data = issue2.json()


issue2_id = get_field(
    issue2_data,
    "id"
)


assert_field(
    issue2_data,
    "epic_id",
    epic_id,
    "Issue epic relationship"
)



# =====================================================
# VERIFY BACKLOG
# =====================================================

backlog = requests.get(
    f"{BASE_URL}/projects/{project_id}/backlog"
)

check(backlog, "Get Backlog")


backlog_data = backlog.json()


if len(backlog_data) != 2:

    print(
        "FAILED: Expected 2 backlog issues"
    )

    sys.exit(1)


titles = [
    get_field(item, "title")
    for item in backlog_data
]


if "Build Login Screen" not in titles:

    print(
        "FAILED: Missing Build Login Screen"
    )

    sys.exit(1)


if "Password Reset" not in titles:

    print(
        "FAILED: Missing Password Reset"
    )

    sys.exit(1)


print("Verified backlog contents")



# =====================================================
# CREATE SPRINT
# =====================================================

sprint = requests.post(
    f"{BASE_URL}/projects/{project_id}/sprints",
    json={
        "name": "Sprint 1",
        "goal": "Authentication release"
    }
)

check(sprint, "Create Sprint")


sprint_data = sprint.json()


sprint_id = get_field(
    sprint_data,
    "id"
)


assert_field(
    sprint_data,
    "name",
    "Sprint 1",
    "Sprint name"
)


assert_field(
    sprint_data,
    "goal",
    "Authentication release",
    "Sprint goal"
)



# =====================================================
# MOVE ISSUE INTO SPRINT
# =====================================================

move = requests.patch(
    f"{BASE_URL}/issues/{issue_id}/sprint",
    json={
        "sprint_id": sprint_id
    }
)

check(move, "Move Issue To Sprint")



# =====================================================
# MOVE ISSUE BACK TO BACKLOG
# =====================================================

move_back = requests.patch(
    f"{BASE_URL}/issues/{issue_id}/sprint",
    json={
        "sprint_id": None
    }
)

check(move_back, "Move Issue Back To Backlog")



# =====================================================
# UPDATE STATUS
# =====================================================

status = requests.patch(
    f"{BASE_URL}/issues/{issue_id}/status",
    json={
        "status": "in_progress"
    }
)

check(status, "Update Issue Status")


verify_status = requests.get(
    f"{BASE_URL}/issues/{issue_id}"
)

check(
    verify_status,
    "Verify Status"
)


assert_field(
    verify_status.json(),
    "status",
    "in_progress",
    "Issue status saved"
)



# =====================================================
# UPDATE STORY POINTS
# =====================================================

points = requests.patch(
    f"{BASE_URL}/issues/{issue_id}/story-points",
    json={
        "story_points": 8
    }
)

check(points, "Update Story Points")


verify_points = requests.get(
    f"{BASE_URL}/issues/{issue_id}"
)

check(
    verify_points,
    "Verify Story Points"
)


assert_field(
    verify_points.json(),
    "story_points",
    8,
    "Story points saved"
)



# =====================================================
# ASSIGN ISSUE TO EPIC
# =====================================================

assign_epic = requests.patch(
    f"{BASE_URL}/issues/{issue_id}/epic",
    json={
        "epic_id": epic_id
    }
)

check(assign_epic, "Assign Issue To Epic")



# =====================================================
# REMOVE ISSUE FROM EPIC
# =====================================================

remove_epic = requests.patch(
    f"{BASE_URL}/issues/{issue_id}/epic",
    json={
        "epic_id": None
    }
)

check(remove_epic, "Remove Issue From Epic")



# =====================================================
# FINAL BACKLOG VERIFY
# =====================================================

final_backlog = requests.get(
    f"{BASE_URL}/projects/{project_id}/backlog"
)

check(
    final_backlog,
    "Final Backlog Check"
)


for item in final_backlog.json():

    sprint_value = get_field(
        item,
        "sprint_id"
    )

    if sprint_value is not None:

        print(
            "FAILED: Backlog issue has sprint assigned"
        )

        sys.exit(1)


print(
    "Verified backlog contains only unscheduled issues"
)



print("\n")
print("==============================")
print(" ALL SYSTEM TESTS PASSED ")
print("==============================")
