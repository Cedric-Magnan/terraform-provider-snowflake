// Code generated by dto builder generator; DO NOT EDIT.

package sdk

import ()

func NewCreateSessionPolicyRequest(
	name SchemaObjectIdentifier,
) *CreateSessionPolicyRequest {
	s := CreateSessionPolicyRequest{}
	s.name = name
	return &s
}

func (s *CreateSessionPolicyRequest) WithOrReplace(OrReplace *bool) *CreateSessionPolicyRequest {
	s.OrReplace = OrReplace
	return s
}

func (s *CreateSessionPolicyRequest) WithIfNotExists(IfNotExists *bool) *CreateSessionPolicyRequest {
	s.IfNotExists = IfNotExists
	return s
}

func (s *CreateSessionPolicyRequest) WithSessionIdleTimeoutMins(SessionIdleTimeoutMins *int) *CreateSessionPolicyRequest {
	s.SessionIdleTimeoutMins = SessionIdleTimeoutMins
	return s
}

func (s *CreateSessionPolicyRequest) WithSessionUiIdleTimeoutMins(SessionUiIdleTimeoutMins *int) *CreateSessionPolicyRequest {
	s.SessionUiIdleTimeoutMins = SessionUiIdleTimeoutMins
	return s
}

func (s *CreateSessionPolicyRequest) WithComment(Comment *string) *CreateSessionPolicyRequest {
	s.Comment = Comment
	return s
}

func NewAlterSessionPolicyRequest(
	name SchemaObjectIdentifier,
) *AlterSessionPolicyRequest {
	s := AlterSessionPolicyRequest{}
	s.name = name
	return &s
}

func (s *AlterSessionPolicyRequest) WithIfExists(IfExists *bool) *AlterSessionPolicyRequest {
	s.IfExists = IfExists
	return s
}

func (s *AlterSessionPolicyRequest) WithRenameTo(RenameTo *SchemaObjectIdentifier) *AlterSessionPolicyRequest {
	s.RenameTo = RenameTo
	return s
}

func (s *AlterSessionPolicyRequest) WithSet(Set *SessionPolicySetRequest) *AlterSessionPolicyRequest {
	s.Set = Set
	return s
}

func (s *AlterSessionPolicyRequest) WithSetTags(SetTags []TagAssociation) *AlterSessionPolicyRequest {
	s.SetTags = SetTags
	return s
}

func (s *AlterSessionPolicyRequest) WithUnsetTags(UnsetTags []ObjectIdentifier) *AlterSessionPolicyRequest {
	s.UnsetTags = UnsetTags
	return s
}

func (s *AlterSessionPolicyRequest) WithUnset(Unset *SessionPolicyUnsetRequest) *AlterSessionPolicyRequest {
	s.Unset = Unset
	return s
}

func NewSessionPolicySetRequest() *SessionPolicySetRequest {
	return &SessionPolicySetRequest{}
}

func (s *SessionPolicySetRequest) WithSessionIdleTimeoutMins(SessionIdleTimeoutMins *int) *SessionPolicySetRequest {
	s.SessionIdleTimeoutMins = SessionIdleTimeoutMins
	return s
}

func (s *SessionPolicySetRequest) WithSessionUiIdleTimeoutMins(SessionUiIdleTimeoutMins *int) *SessionPolicySetRequest {
	s.SessionUiIdleTimeoutMins = SessionUiIdleTimeoutMins
	return s
}

func (s *SessionPolicySetRequest) WithComment(Comment *string) *SessionPolicySetRequest {
	s.Comment = Comment
	return s
}

func NewSessionPolicyUnsetRequest() *SessionPolicyUnsetRequest {
	return &SessionPolicyUnsetRequest{}
}

func (s *SessionPolicyUnsetRequest) WithSessionIdleTimeoutMins(SessionIdleTimeoutMins *bool) *SessionPolicyUnsetRequest {
	s.SessionIdleTimeoutMins = SessionIdleTimeoutMins
	return s
}

func (s *SessionPolicyUnsetRequest) WithSessionUiIdleTimeoutMins(SessionUiIdleTimeoutMins *bool) *SessionPolicyUnsetRequest {
	s.SessionUiIdleTimeoutMins = SessionUiIdleTimeoutMins
	return s
}

func (s *SessionPolicyUnsetRequest) WithComment(Comment *bool) *SessionPolicyUnsetRequest {
	s.Comment = Comment
	return s
}

func NewDropSessionPolicyRequest(
	name SchemaObjectIdentifier,
) *DropSessionPolicyRequest {
	s := DropSessionPolicyRequest{}
	s.name = name
	return &s
}

func (s *DropSessionPolicyRequest) WithIfExists(IfExists *bool) *DropSessionPolicyRequest {
	s.IfExists = IfExists
	return s
}

func NewShowSessionPolicyRequest() *ShowSessionPolicyRequest {
	return &ShowSessionPolicyRequest{}
}

func NewDescribeSessionPolicyRequest(
	name SchemaObjectIdentifier,
) *DescribeSessionPolicyRequest {
	s := DescribeSessionPolicyRequest{}
	s.name = name
	return &s
}
