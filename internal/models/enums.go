package models

type IssueType string

const (
	Story  IssueType = "story"
	Task   IssueType = "task"
	Bug    IssueType = "bug"
	Subtask IssueType = "subtask"
)

type IssueStatus string

const (
	Todo       IssueStatus = "todo"
	InProgress IssueStatus = "in_progress"
	Review     IssueStatus = "review"
	Done       IssueStatus = "done"
	Blocked    IssueStatus = "blocked"
)

type Priority string

const (
	Lowest  Priority = "lowest"
	Low     Priority = "low"
	Medium  Priority = "medium"
	High    Priority = "high"
	Highest Priority = "highest"
)
