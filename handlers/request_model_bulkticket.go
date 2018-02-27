package handlers
// The following Type defines the Input JSON Body for Creating a Bulk Ticket Request
type BulkTicket_Request struct{
	UserLogin	string	`json:"UserLogin,omitempty"`
	Password	string	`json:"Password,omitempty"`
	Action		string	`json:"Action,omitempty"`
	TicketIDs	[]string	`json:"TicketIDs,omitempty"`
	MergeTo		string	`json:"MergeTo,omitempty"`
	Subaction	string	`json:"Subaction,omitempty"`
	Priority	string	`json:"Priority,omitempty"`
	TypeID		int	`json:"TypeID,omitempty"`
	StateID		string	`json:"StateID,omitempty"`
	QueueID		int	`json:"QueueID,omitempty"`
	Owner		string	`json:"Owner,omitempty"`
	ResponsibleID	string	`json:"ResponsibleID,omitempty"`
	Responsible	string	`json:"Responsible,omitempty"`
	PriorityID	string	`json:"PriorityID,omitempty"`
	Queue		string	`json:"Queue,omitempty"`
	Subject		string	`json:"Subject,omitempty"`
	Body		string	`json:"Body,omitempty"`
	ArticleTypeID	string	`json:"ArticleTypeID,omitempty"`
	ArticleType	string	`json:"ArticleType,omitempty"`
	State		string	`json:"State,omitempty"`
	MergeToSelection	string	`json:"MergeToSelection,omitempty"`
	LinkTogether	string	`json:"LinkTogether,omitempty"`
	EmailSubject	string	`json:"EmailSubject,omitempty"`
	EmailBody	string	`json:"EmailBody,omitempty"`
	EmailTimeUnits	string	`json:"EmailTimeUnits,omitempty"`
	LinkTogetherParent	string	`json:"LinkTogetherParent,omitempty"`
	Unlock		string	`json:"Unlock,omitempty"`
	MergeToChecked	string	`json:"MergeToChecked,omitempty"`
	MergeToOldestChecked	string	`json:"MergeToOldestChecked,omitempty"`
	Year	string	`json:"Year,omitempty"`
	Month	string	`json:"Month,omitempty"`
	Day	string	`json:"Day,omitempty"`
	Hour	string	`json:"Hour,omitempty"`
	Minute	string	`json:"Minute,omitempty"`
        OwnerID	string	`json:"OwnerID,omitempty"`
}
