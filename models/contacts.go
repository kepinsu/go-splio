package models

type ListSearchContactsQuery struct {
	PerPage    int
	PageNumber string
	Fields     []ListSearchContactsField
}

type ListSearchContactsField struct {
	Key      string
	Operator string
	Value    string
}

type ListSearchContacts struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []Contact `json:"elements,omitempty"`
}

type Contact struct {
	ID           uint               `json:"id,omitempty"`
	Lastname     string             `json:"lastname,omitempty"`
	Firstname    string             `json:"firstname,omitempty"`
	Email        string             `json:"email,omitempty"`
	Cellphone    string             `json:"cellphone,omitempty"`
	CreationDate string             `json:"creation_date,omitempty"`
	Language     string             `json:"language,omitempty"`
	CustomFields []CustomField      `json:"custom_fields,omitempty"`
	List         []ListContactField `json:"list,omitempty"`
	Loyalty      []Loyalty          `json:"loyalty,omitempty"`
}

type ListContactField struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CustomField struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Value    string `json:"value,omitempty"`
	DataType string `json:"data_type,omitempty"`
}

type Loyalty struct {
	CardCode  string `json:"card_code,omitempty"`
	IdProgram string `json:"id_program,omitempty"`
}

type CreateContact struct {
	Lastname     string             `json:"lastname,omitempty"`
	Firstname    string             `json:"firstname,omitempty"`
	Email        string             `json:"email,omitempty"`
	Cellphone    string             `json:"cellphone,omitempty"`
	CreationDate string             `json:"creation_date,omitempty"`
	Language     string             `json:"language,omitempty"`
	DoubleOptin  DoubleOptin        `json:"double_optin,omitempty"`
	CustomFields []CustomField      `json:"custom_fields,omitempty"`
	List         []ListContactField `json:"list,omitempty"`
	Loyalty      []Loyalty          `json:"loyalty,omitempty"`
}

type DoubleOptin struct {
	Message       string `json:"message,omitempty"`
	Reminder      string `json:"reminder,omitempty"`
	ReminderDelay int    `json:"reminder_delay,omitempty"`
}

type BulkDeleteContact struct {
	Contacts []string `json:"contacts,omitempty"`
}

type ContactActivities struct {
	CountElement uint                       `json:"count_element,omitempty"`
	CuurentPage  uint                       `json:"cuurent_page,omitempty"`
	PerPage      uint                       `json:"per_page,omitempty"`
	Elements     []ContactActivitiesElement `json:"elements,omitempty"`
}

type ContactActivitiesElement struct {
	Channel            string  `json:"channel,omitempty"`
	Contact            *string `json:"contact,omitempty"`
	ChannelDetailsID   string  `json:"channel_details_id,omitempty"`
	ChannelDetailsName string  `json:"channel_details_name,omitempty"`
	EventType          string  `json:"event_type,omitempty"`
	Event              string  `json:"event,omitempty"`
	ValueDate          string  `json:"value_date,omitempty"`
	SourceID           string  `json:"source_id,omitempty"`
	SourceName         string  `json:"source_name,omitempty"`
	SourceType         string  `json:"source_type,omitempty"`
	SourceDate         string  `json:"source_date,omitempty"`
}

type OptOutContactRequest struct {
	Contacts []string `json:"contacts,omitempty"`
}

type OptOutContactResponse struct {
	Time   int                         `json:"time,omitempty"`
	Errors int                         `json:"errors,omitempty"`
	Items  []OptOutContactResponseItem `json:"items,omitempty"`
}

type OptOutContactResponseItem struct {
	ContactID   string `json:"contact_id,omitempty"`
	Code        int    `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
}

type ListMemberShipLists struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []MemberShip `json:"elements,omitempty"`
}

type MemberShip struct {
	Id       int    `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Members  uint   `json:"members,omitempty"`
	NbEmails uint   `json:"nb_emails,omitempty"`
	NbSms    uint   `json:"nb_sms,omitempty"`
}

type SubscribeContactToListRequest struct {
	ListIDs []int `json:"list_ids,omitempty"`
}

type SubscribeContactToListResponse struct {
	Lastname     string             `json:"lastname,omitempty"`
	Firstname    string             `json:"firstname,omitempty"`
	Email        string             `json:"email,omitempty"`
	Cellphone    string             `json:"cellphone,omitempty"`
	CreationDate string             `json:"creation_date,omitempty"`
	Language     string             `json:"language,omitempty"`
	CustomFields []CustomField      `json:"custom_fields,omitempty"`
	List         []ListContactField `json:"list,omitempty"`
}

type UnsubscribeContactToListRequest struct {
	ListIDs []int `json:"list_ids,omitempty"`
}

type UnsubscribeContactToListResponse struct {
	Lastname     string             `json:"lastname,omitempty"`
	Firstname    string             `json:"firstname,omitempty"`
	Email        string             `json:"email,omitempty"`
	Cellphone    string             `json:"cellphone,omitempty"`
	CreationDate string             `json:"creation_date,omitempty"`
	Language     string             `json:"language,omitempty"`
	CustomFields []CustomField      `json:"custom_fields,omitempty"`
	List         []ListContactField `json:"list,omitempty"`
}

type ContactsAbandonnedCarts struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []ContactOrder `json:"elements,omitempty"`
}

type ContactsOrders struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []ContactOrder `json:"elements,omitempty"`
}

type ContactOrder struct {
	ExternalID     int         `json:"external_id,omitempty"`
	CreatedAt      string      `json:"created_at,omitempty"`
	OrderedAt      string      `json:"ordered_at,omitempty"`
	ShippingAmount float32     `json:"shipping_amount,omitempty"`
	DiscountAmount float32     `json:"discount_amount,omitempty"`
	TotalPrice     float32     `json:"total_price,omitempty"`
	TaxAmount      float32     `json:"tax_amount,omitempty"`
	Currency       string      `json:"currency,omitempty"`
	SalesPerson    string      `json:"sales_person,omitempty"`
	Store          *OrderStore `json:"store,omitempty"`
	CustomFields   []struct {
		ID    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"custom_fields,omitempty"`
}

type OrderStore struct {
	ExternalID int    `json:"external_id,omitempty"`
	Name       string `json:"name,omitempty"`
}

type Product struct {
	ExternalID   string `json:"external_id,omitempty"`
	Description  string `json:"description,omitempty"`
	Brand        string `json:"brand,omitempty"`
	Category     string `json:"category,omitempty"`
	Price        uint   `json:"price,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	CustomFields []struct {
		ID       int    `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Value    string `json:"value,omitempty"`
		DataType string `json:"data_type,omitempty"`
	} `json:"custom_fields,omitempty"`
}

type ContactsStore struct {
	CountElement uint `json:"count_element,omitempty"`
	CuurentPage  uint `json:"cuurent_page,omitempty"`
	PerPage      uint `json:"per_page,omitempty"`
	Sort         struct {
		ID string `json:"id,omitempty"`
	} `json:"sort,omitempty"`
	Elements []Store `json:"elements,omitempty"`
}

type Store struct {
	ExternalID   string `json:"external_id,omitempty"`
	Name         string `json:"name,omitempty"`
	Online       bool
	StoreType    string
	Manager      string
	CreatedAt    string `json:"created_at,omitempty"`
	UpdatedAt    string `json:"updated_at,omitempty"`
	CustomFields []struct {
		ID       int    `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Value    string `json:"value,omitempty"`
		DataType string `json:"data_type,omitempty"`
	} `json:"custom_fields,omitempty"`
}
