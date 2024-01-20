
// Code generated by Script. DO NOT EDIT.
// Source: script/codegen/main.go
//
// Generated by this command:
//
//	go run script/codegen/main.go

package zendesk

import "context"

func (z *Client) GetTicketFormsIterator(ctx context.Context, opts *PaginationOptions) *Iterator[TicketForm] {
	return &Iterator[TicketForm]{
		CommonOptions: opts.CommonOptions,
		pageSize:      opts.PageSize,
		hasMore:       true,
		isCBP:         opts.IsCBP,
		pageAfter:     "",
		pageIndex:     1,
		ctx:           ctx,
		obpFunc:       z.GetTicketFormsOBP,
		cbpFunc:       z.GetTicketFormsCBP,
	}
}

func (z *Client) GetTicketFormsOBP(ctx context.Context, opts *OBPOptions) ([]TicketForm, Page, error) {
	var data struct {
		TicketForms []TicketForm `json:"ticket_forms"`
		Page
	}

	tmp := opts
	if tmp == nil {
		tmp = &OBPOptions{}
	}
	
	u, err := addOptions("/ticket_forms.json", tmp)
	
	if err != nil {
		return nil, Page{}, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, Page{}, err
	}
	return data.TicketForms, data.Page, nil
}

func (z *Client) GetTicketFormsCBP(ctx context.Context, opts *CBPOptions) ([]TicketForm, CursorPaginationMeta, error) {
	var data struct {
		TicketForms []TicketForm `json:"ticket_forms"`
		Meta    CursorPaginationMeta `json:"meta"`
	}

	tmp := opts
	if tmp == nil {
		tmp = &CBPOptions{}
	}
	
	u, err := addOptions("/ticket_forms.json", tmp)
	
	if err != nil {
		return nil, data.Meta, err
	}

	err = getData(z, ctx, u, &data)
	if err != nil {
		return nil, data.Meta, err
	}
	return data.TicketForms, data.Meta, nil
}

