package telegram

import (
	"log"
	"sync"
)

type (
	Filters struct {
		IsPrivate      bool
		IsGroup        bool
		IsChannel      bool
		IsCommand      bool
		IsText         bool
		IsMedia        bool
		Func           func(*NewMessage) bool
		BlackListChats []int64
		WhiteListChats []int64
		Users          []int64
		Outgoing       bool
		Incoming       bool
	}

	MessageHandle struct {
		Pattern interface{}
		Handler func(m *NewMessage) error
		Client  *Client
		Filters *Filters
	}

	ChatActionHandle struct {
		Handler func(m *NewMessage) error
		Client  *Client
	}

	InlineHandle struct {
		Pattern interface{}
		Handler func(m *InlineQuery) error
		Client  *Client
	}

	CallbackHandle struct {
		Pattern interface{}
		Handler func(m *CallbackQuery) error
		Client  *Client
	}

	RawHandle struct {
		updateType Update
		Handler    func(m Update) error
		Client     *Client
	}

	MessageEditHandle struct {
		Pattern interface{}
		Handler func(m *NewMessage) error
		Client  *Client
	}

	Command struct {
		Cmd    string
		Prefix string
	}
)

type (
	Progress struct {
		current int64
		total   int64
		rwlock  sync.RWMutex
	}

	DialogOptions struct {
		OffsetID      int32     `json:"offset_id,omitempty"`
		OffsetDate    int32     `json:"offset_date,omitempty"`
		OffsetPeer    InputPeer `json:"offset_peer,omitempty"`
		Limit         int32     `json:"limit,omitempty"`
		ExcludePinned bool      `json:"exclude_pinned,omitempty"`
		FolderID      int32     `json:"folder_id,omitempty"`
	}

	InlineOptions struct {
		Dialog   interface{}
		Offset   int32
		Query    string
		GeoPoint InputGeoPoint
	}

	LoginOptions struct {
		Password  string `json:"password,omitempty"`
		Code      string `json:"code,omitempty"`
		CodeHash  string `json:"code_hash,omitempty"`
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
	}

	DownloadOptions struct {
		Progress *Progress `json:"progress,omitempty"`
		FileName string    `json:"file_name,omitempty"`
		DcID     int32     `json:"dc_id,omitempty"`
		Size     int32     `json:"size,omitempty"`
	}

	SendOptions struct {
		ReplyID       int32               `json:"reply_id,omitempty"`
		Caption       interface{}         `json:"caption,omitempty"`
		ParseMode     string              `json:"parse_mode,omitempty"`
		Silent        bool                `json:"silent,omitempty"`
		LinkPreview   bool                `json:"link_preview,omitempty"`
		ReplyMarkup   ReplyMarkup         `json:"reply_markup,omitempty"`
		ClearDraft    bool                `json:"clear_draft,omitempty"`
		NoForwards    bool                `json:"no_forwards,omitempty"`
		ScheduleDate  int32               `json:"schedule_date,omitempty"`
		SendAs        InputPeer           `json:"send_as,omitempty"`
		Thumb         InputFile           `json:"thumb,omitempty"`
		TTL           int32               `json:"ttl,omitempty"`
		ForceDocument bool                `json:"force_document,omitempty"`
		FileName      string              `json:"file_name,omitempty"`
		Attributes    []DocumentAttribute `json:"attributes,omitempty"`
		Media         interface{}
	}

	MediaOptions struct {
		Caption       interface{}         `json:"caption,omitempty"`
		ParseMode     string              `json:"parse_mode,omitempty"`
		Silent        bool                `json:"silent,omitempty"`
		LinkPreview   bool                `json:"link_preview,omitempty"`
		ReplyMarkup   ReplyMarkup         `json:"reply_markup,omitempty"`
		ClearDraft    bool                `json:"clear_draft,omitempty"`
		NoForwards    bool                `json:"no_forwards,omitempty"`
		Thumb         InputFile           `json:"thumb,omitempty"`
		NoSoundVideo  bool                `json:"no_sound_video,omitempty"`
		ForceDocument bool                `json:"force_document,omitempty"`
		ReplyID       int32               `json:"reply_id,omitempty"`
		FileName      string              `json:"file_name,omitempty"`
		TTL           int32               `json:"ttl,omitempty"`
		Attributes    []DocumentAttribute `json:"attributes,omitempty"`
		ScheduleDate  int32               `json:"schedule_date,omitempty"`
		SendAs        InputPeer           `json:"send_as,omitempty"`
	}

	CustomAttrs struct {
		FileName      string              `json:"file_name,omitempty"`
		Thumb         InputFile           `json:"thumb,omitempty"`
		Attributes    []DocumentAttribute `json:"attributes,omitempty"`
		ForceDocument bool                `json:"force_document,omitempty"`
		TTL           int32               `json:"ttl,omitempty"`
	}

	ForwardOptions struct {
		HideCaption bool `json:"hide_caption,omitempty"`
		HideAuthor  bool `json:"hide_author,omitempty"`
		Silent      bool `json:"silent,omitempty"`
		Protected   bool `json:"protected,omitempty"`
	}

	Participant struct {
		User        *UserObj           `json:"user,omitempty"`
		Admin       bool               `json:"admin,omitempty"`
		Banned      bool               `json:"banned,omitempty"`
		Creator     bool               `json:"creator,omitempty"`
		Left        bool               `json:"left,omitempty"`
		Participant ChannelParticipant `json:"participant,omitempty"`
		Rights      *ChatAdminRights   `json:"rights,omitempty"`
	}

	ParticipantOptions struct {
		Query  string                    `json:"query,omitempty"`
		Filter ChannelParticipantsFilter `json:"filter,omitempty"`
		Offset int32                     `json:"offset,omitempty"`
		Limit  int32                     `json:"limit,omitempty"`
	}

	AdminOptions struct {
		IsAdmin bool             `json:"is_admin,omitempty"`
		Rights  *ChatAdminRights `json:"rights,omitempty"`
		Rank    string           `json:"rank,omitempty"`
	}

	PhotosOptions struct {
		MaxID  int64 `json:"max_id,omitempty"`
		Offset int32 `json:"offset,omitempty"`
		Limit  int32 `json:"limit,omitempty"`
	}

	PinOptions struct {
		Unpin     bool `json:"unpin,omitempty"`
		PmOneside bool `json:"pm_oneside,omitempty"`
		Silent    bool `json:"silent,omitempty"`
	}

	BannedOptions struct {
		Ban    bool              `json:"ban,omitempty"`
		Unban  bool              `json:"unban,omitempty"`
		Mute   bool              `json:"mute,omitempty"`
		Unmute bool              `json:"unmute,omitempty"`
		Rights *ChatBannedRights `json:"rights,omitempty"`
	}

	ActionResult struct {
		Peer   InputPeer `json:"peer,omitempty"`
		Client *Client   `json:"client,omitempty"`
	}

	InlineSendOptions struct {
		Gallery      bool   `json:"gallery,omitempty"`
		NextOffset   string `json:"next_offset,omitempty"`
		CacheTime    int32  `json:"cache_time,omitempty"`
		Private      bool   `json:"private,omitempty"`
		SwitchPm     string `json:"switch_pm,omitempty"`
		SwitchPmText string `json:"switch_pm_text,omitempty"`
	}

	CallbackOptions struct {
		Alert     bool   `json:"alert,omitempty"`
		CacheTime int32  `json:"cache_time,omitempty"`
		URL       string `json:"url,omitempty"`
	}

	ArticleOptions struct {
		ID           string                             `json:"id,omitempty"`
		ExcludeMedia bool                               `json:"exclude_media,omitempty"`
		Thumb        InputWebDocument                   `json:"thumb,omitempty"`
		Content      InputWebDocument                   `json:"content,omitempty"`
		LinkPreview  bool                               `json:"link_preview,omitempty"`
		ReplyMarkup  ReplyMarkup                        `json:"reply_markup,omitempty"`
		Entities     []MessageEntity                    `json:"entities,omitempty"`
		ParseMode    string                             `json:"parse_mode,omitempty"`
		Caption      string                             `json:"caption,omitempty"`
		Venue        *InputBotInlineMessageMediaVenue   `json:"venue,omitempty"`
		Location     *InputBotInlineMessageMediaGeo     `json:"location,omitempty"`
		Contact      *InputBotInlineMessageMediaContact `json:"contact,omitempty"`
		Invoice      *InputBotInlineMessageMediaInvoice `json:"invoice,omitempty"`
	}

	SearchOption struct {
		IDs      []int32        `json:"ids,omitempty"`
		Query    string         `json:"query,omitempty"`
		Offset   int32          `json:"offset,omitempty"`
		Limit    int32          `json:"limit,omitempty"`
		Filter   MessagesFilter `json:"filter,omitempty"`
		TopMsgID int32          `json:"top_msg_id,omitempty"`
		MaxID    int32          `json:"max_id,omitempty"`
		MinID    int32          `json:"min_id,omitempty"`
		MaxDate  int32          `json:"max_date,omitempty"`
		MinDate  int32          `json:"min_date,omitempty"`
	}

	TitleOptions struct {
		LastName string `json:"last_name,omitempty"`
		About    string `json:"about,omitempty"`
	}

	Log struct {
		Logger *log.Logger
	}
)

var (
	ParticipantsAdmins = &ParticipantOptions{
		Filter: &ChannelParticipantsAdmins{},
		Query:  "",
		Offset: 0,
		Limit:  50,
	}
)

func (p *Progress) Set(value int64) {
	p.rwlock.Lock()
	defer p.rwlock.Unlock()
	p.current = value
}

func (p *Progress) Get() int64 {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()
	return p.current
}

func (p *Progress) Data() (total int64, current int64) {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()
	return p.total, p.current
}

func (p *Progress) Percentage() float64 {
	p.rwlock.RLock()
	defer p.rwlock.RUnlock()
	return float64(p.current) / float64(p.total) * 100
}

// Cancel the pointed Action,
// Returns true if the action was cancelled
func (a *ActionResult) Cancel() bool {
	if a.Peer == nil || a.Client == nil {
		return false // Avoid nil pointer dereference
	}
	b, err := a.Client.MessagesSetTyping(a.Peer, 0, &SendMessageCancelAction{})
	if err != nil {
		return false
	}
	return b
}

// ChatMember Rights
func (p Participant) IsCreator() bool {
	return p.Creator
}

func (p Participant) IsAdmin() bool {
	return p.Admin
}

func (p Participant) IsBanned() bool {
	return p.Banned
}

func (p Participant) IsLeft() bool {
	return p.Left
}

func (p Participant) GetUser() *UserObj {
	return p.User
}

func (p Participant) GetParticipant() ChannelParticipant {
	return p.Participant
}

func (r ChatAdminRights) CanChangeInfo() bool {
	return r.ChangeInfo
}

func (r ChatAdminRights) CanPostMessages() bool {
	return r.PostMessages
}

func (r ChatAdminRights) CanEditMessages() bool {
	return r.EditMessages
}

func (r ChatAdminRights) CanDeleteMessages() bool {
	return r.DeleteMessages
}

func (r ChatAdminRights) CanBanUsers() bool {
	return r.BanUsers
}

func (r ChatAdminRights) CanInviteUsers() bool {
	return r.InviteUsers
}

func (r ChatAdminRights) CanPinMessages() bool {
	return r.PinMessages
}

func (r ChatAdminRights) CanPromoteMembers() bool {
	return r.AddAdmins
}

func (r ChatAdminRights) IsAnonymous() bool {
	return r.Anonymous
}

func (r ChatAdminRights) CanManageCall() bool {
	return r.ManageCall
}

func (p Participant) GetRank() string {
	if pp, ok := p.Participant.(*ChannelParticipantCreator); ok {
		return pp.Rank
	} else if pp, ok := p.Participant.(*ChannelParticipantAdmin); ok {
		return pp.Rank
	}
	return ""
}

//Custom Logger

func (l *Log) Error(e error) {
	l.Logger.Println("Client - Error -", e)
}
