// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feedfilter/feedfilter_.proto

/*
Package feedfilter is a generated protocol buffer package.

It is generated from these files:
	feedfilter/feedfilter_.proto

It has these top-level messages:
	FeedFilter
	PostFilter
	SocialAccountFilter
*/
package feedfilter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import feedkindfilter "github.com/21stio/go-playground/grpc/schema/feedkindfilter"
import infofilter "github.com/21stio/go-playground/grpc/schema/infofilter"
import intfilter "github.com/21stio/go-playground/grpc/schema/intfilter"
import idfilter "github.com/21stio/go-playground/grpc/schema/idfilter"
import servicefilter "github.com/21stio/go-playground/grpc/schema/servicefilter"
import textfilter "github.com/21stio/go-playground/grpc/schema/textfilter"
import liketermkindfilter "github.com/21stio/go-playground/grpc/schema/liketermkindfilter"
import commenttermkindfilter "github.com/21stio/go-playground/grpc/schema/commenttermkindfilter"
import sharetermkindfilter "github.com/21stio/go-playground/grpc/schema/sharetermkindfilter"
import mediafilter "github.com/21stio/go-playground/grpc/schema/mediafilter"
import socialaccountkindfilter "github.com/21stio/go-playground/grpc/schema/socialaccountkindfilter"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FeedFilter struct {
	Kind             *feedkindfilter.FeedKindFilter `protobuf:"bytes,1,opt,name=kind" json:"kind,omitempty"`
	Info             *infofilter.InfoFilter         `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
	CreatedBy        *SocialAccountFilter           `protobuf:"bytes,3,opt,name=createdBy" json:"createdBy,omitempty"`
	Count            *intfilter.IntFilter           `protobuf:"bytes,4,opt,name=count" json:"count,omitempty"`
	PostsSome        *PostFilter                    `protobuf:"bytes,5,opt,name=postsSome" json:"postsSome,omitempty"`
	PostsEvery       *PostFilter                    `protobuf:"bytes,6,opt,name=postsEvery" json:"postsEvery,omitempty"`
	PostsNone        *PostFilter                    `protobuf:"bytes,7,opt,name=postsNone" json:"postsNone,omitempty"`
	IdsSome          *idfilter.IdFilter             `protobuf:"bytes,8,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery         *idfilter.IdFilter             `protobuf:"bytes,9,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone          *idfilter.IdFilter             `protobuf:"bytes,10,opt,name=idsNone" json:"idsNone,omitempty"`
	Meta             *servicefilter.TypeMetaFilter  `protobuf:"bytes,11,opt,name=meta" json:"meta,omitempty"`
	And              []*FeedFilter                  `protobuf:"bytes,12,rep,name=and" json:"and,omitempty"`
	Or               []*FeedFilter                  `protobuf:"bytes,13,rep,name=or" json:"or,omitempty"`
	Not              []*FeedFilter                  `protobuf:"bytes,14,rep,name=not" json:"not,omitempty"`
	Hash             *string                        `protobuf:"bytes,15,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                         `json:"-"`
}

func (m *FeedFilter) Reset()                    { *m = FeedFilter{} }
func (m *FeedFilter) String() string            { return proto.CompactTextString(m) }
func (*FeedFilter) ProtoMessage()               {}
func (*FeedFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FeedFilter) GetKind() *feedkindfilter.FeedKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *FeedFilter) GetInfo() *infofilter.InfoFilter {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *FeedFilter) GetCreatedBy() *SocialAccountFilter {
	if m != nil {
		return m.CreatedBy
	}
	return nil
}

func (m *FeedFilter) GetCount() *intfilter.IntFilter {
	if m != nil {
		return m.Count
	}
	return nil
}

func (m *FeedFilter) GetPostsSome() *PostFilter {
	if m != nil {
		return m.PostsSome
	}
	return nil
}

func (m *FeedFilter) GetPostsEvery() *PostFilter {
	if m != nil {
		return m.PostsEvery
	}
	return nil
}

func (m *FeedFilter) GetPostsNone() *PostFilter {
	if m != nil {
		return m.PostsNone
	}
	return nil
}

func (m *FeedFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *FeedFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *FeedFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *FeedFilter) GetMeta() *servicefilter.TypeMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *FeedFilter) GetAnd() []*FeedFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *FeedFilter) GetOr() []*FeedFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *FeedFilter) GetNot() []*FeedFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *FeedFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

type PostFilter struct {
	Author           *SocialAccountFilter                         `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
	Content          *textfilter.TextFilter                       `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	LikeTerm         *liketermkindfilter.LikeTermKindFilter       `protobuf:"bytes,3,opt,name=likeTerm" json:"likeTerm,omitempty"`
	LikesCount       *intfilter.IntFilter                         `protobuf:"bytes,4,opt,name=likesCount" json:"likesCount,omitempty"`
	LikedBySome      *SocialAccountFilter                         `protobuf:"bytes,5,opt,name=likedBySome" json:"likedBySome,omitempty"`
	LikedByEvery     *SocialAccountFilter                         `protobuf:"bytes,6,opt,name=likedByEvery" json:"likedByEvery,omitempty"`
	LikedByNone      *SocialAccountFilter                         `protobuf:"bytes,7,opt,name=likedByNone" json:"likedByNone,omitempty"`
	CommentTerm      *commenttermkindfilter.CommentTermKindFilter `protobuf:"bytes,8,opt,name=commentTerm" json:"commentTerm,omitempty"`
	CommentsCount    *intfilter.IntFilter                         `protobuf:"bytes,9,opt,name=commentsCount" json:"commentsCount,omitempty"`
	CommentedBySome  *SocialAccountFilter                         `protobuf:"bytes,10,opt,name=commentedBySome" json:"commentedBySome,omitempty"`
	CommentedByEvery *SocialAccountFilter                         `protobuf:"bytes,11,opt,name=commentedByEvery" json:"commentedByEvery,omitempty"`
	CommentedByNone  *SocialAccountFilter                         `protobuf:"bytes,12,opt,name=commentedByNone" json:"commentedByNone,omitempty"`
	CommentsSome     *PostFilter                                  `protobuf:"bytes,13,opt,name=commentsSome" json:"commentsSome,omitempty"`
	CommentsEvery    *PostFilter                                  `protobuf:"bytes,14,opt,name=commentsEvery" json:"commentsEvery,omitempty"`
	CommentsNone     *PostFilter                                  `protobuf:"bytes,15,opt,name=commentsNone" json:"commentsNone,omitempty"`
	Commenting       *PostFilter                                  `protobuf:"bytes,16,opt,name=commenting" json:"commenting,omitempty"`
	ShareTerm        *sharetermkindfilter.ShareTermKindFilter     `protobuf:"bytes,17,opt,name=shareTerm" json:"shareTerm,omitempty"`
	SharesCount      *intfilter.IntFilter                         `protobuf:"bytes,18,opt,name=sharesCount" json:"sharesCount,omitempty"`
	Shared           *PostFilter                                  `protobuf:"bytes,19,opt,name=shared" json:"shared,omitempty"`
	SharedBySome     *SocialAccountFilter                         `protobuf:"bytes,20,opt,name=sharedBySome" json:"sharedBySome,omitempty"`
	SharedByEvery    *SocialAccountFilter                         `protobuf:"bytes,21,opt,name=sharedByEvery" json:"sharedByEvery,omitempty"`
	SharedByNone     *SocialAccountFilter                         `protobuf:"bytes,22,opt,name=sharedByNone" json:"sharedByNone,omitempty"`
	AttachmentsSome  *mediafilter.MediaFilter                     `protobuf:"bytes,23,opt,name=attachmentsSome" json:"attachmentsSome,omitempty"`
	AttachmentsEvery *mediafilter.MediaFilter                     `protobuf:"bytes,24,opt,name=attachmentsEvery" json:"attachmentsEvery,omitempty"`
	AttachmentsNone  *mediafilter.MediaFilter                     `protobuf:"bytes,25,opt,name=attachmentsNone" json:"attachmentsNone,omitempty"`
	IdsSome          *idfilter.IdFilter                           `protobuf:"bytes,26,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery         *idfilter.IdFilter                           `protobuf:"bytes,27,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone          *idfilter.IdFilter                           `protobuf:"bytes,28,opt,name=idsNone" json:"idsNone,omitempty"`
	Meta             *servicefilter.TypeMetaFilter                `protobuf:"bytes,29,opt,name=meta" json:"meta,omitempty"`
	And              []*PostFilter                                `protobuf:"bytes,30,rep,name=and" json:"and,omitempty"`
	Or               []*PostFilter                                `protobuf:"bytes,31,rep,name=or" json:"or,omitempty"`
	Not              []*PostFilter                                `protobuf:"bytes,32,rep,name=not" json:"not,omitempty"`
	Hash             *string                                      `protobuf:"bytes,33,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                       `json:"-"`
}

func (m *PostFilter) Reset()                    { *m = PostFilter{} }
func (m *PostFilter) String() string            { return proto.CompactTextString(m) }
func (*PostFilter) ProtoMessage()               {}
func (*PostFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PostFilter) GetAuthor() *SocialAccountFilter {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *PostFilter) GetContent() *textfilter.TextFilter {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *PostFilter) GetLikeTerm() *liketermkindfilter.LikeTermKindFilter {
	if m != nil {
		return m.LikeTerm
	}
	return nil
}

func (m *PostFilter) GetLikesCount() *intfilter.IntFilter {
	if m != nil {
		return m.LikesCount
	}
	return nil
}

func (m *PostFilter) GetLikedBySome() *SocialAccountFilter {
	if m != nil {
		return m.LikedBySome
	}
	return nil
}

func (m *PostFilter) GetLikedByEvery() *SocialAccountFilter {
	if m != nil {
		return m.LikedByEvery
	}
	return nil
}

func (m *PostFilter) GetLikedByNone() *SocialAccountFilter {
	if m != nil {
		return m.LikedByNone
	}
	return nil
}

func (m *PostFilter) GetCommentTerm() *commenttermkindfilter.CommentTermKindFilter {
	if m != nil {
		return m.CommentTerm
	}
	return nil
}

func (m *PostFilter) GetCommentsCount() *intfilter.IntFilter {
	if m != nil {
		return m.CommentsCount
	}
	return nil
}

func (m *PostFilter) GetCommentedBySome() *SocialAccountFilter {
	if m != nil {
		return m.CommentedBySome
	}
	return nil
}

func (m *PostFilter) GetCommentedByEvery() *SocialAccountFilter {
	if m != nil {
		return m.CommentedByEvery
	}
	return nil
}

func (m *PostFilter) GetCommentedByNone() *SocialAccountFilter {
	if m != nil {
		return m.CommentedByNone
	}
	return nil
}

func (m *PostFilter) GetCommentsSome() *PostFilter {
	if m != nil {
		return m.CommentsSome
	}
	return nil
}

func (m *PostFilter) GetCommentsEvery() *PostFilter {
	if m != nil {
		return m.CommentsEvery
	}
	return nil
}

func (m *PostFilter) GetCommentsNone() *PostFilter {
	if m != nil {
		return m.CommentsNone
	}
	return nil
}

func (m *PostFilter) GetCommenting() *PostFilter {
	if m != nil {
		return m.Commenting
	}
	return nil
}

func (m *PostFilter) GetShareTerm() *sharetermkindfilter.ShareTermKindFilter {
	if m != nil {
		return m.ShareTerm
	}
	return nil
}

func (m *PostFilter) GetSharesCount() *intfilter.IntFilter {
	if m != nil {
		return m.SharesCount
	}
	return nil
}

func (m *PostFilter) GetShared() *PostFilter {
	if m != nil {
		return m.Shared
	}
	return nil
}

func (m *PostFilter) GetSharedBySome() *SocialAccountFilter {
	if m != nil {
		return m.SharedBySome
	}
	return nil
}

func (m *PostFilter) GetSharedByEvery() *SocialAccountFilter {
	if m != nil {
		return m.SharedByEvery
	}
	return nil
}

func (m *PostFilter) GetSharedByNone() *SocialAccountFilter {
	if m != nil {
		return m.SharedByNone
	}
	return nil
}

func (m *PostFilter) GetAttachmentsSome() *mediafilter.MediaFilter {
	if m != nil {
		return m.AttachmentsSome
	}
	return nil
}

func (m *PostFilter) GetAttachmentsEvery() *mediafilter.MediaFilter {
	if m != nil {
		return m.AttachmentsEvery
	}
	return nil
}

func (m *PostFilter) GetAttachmentsNone() *mediafilter.MediaFilter {
	if m != nil {
		return m.AttachmentsNone
	}
	return nil
}

func (m *PostFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *PostFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *PostFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *PostFilter) GetMeta() *servicefilter.TypeMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *PostFilter) GetAnd() []*PostFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *PostFilter) GetOr() []*PostFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *PostFilter) GetNot() []*PostFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *PostFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

type SocialAccountFilter struct {
	Info             *infofilter.InfoFilter                           `protobuf:"bytes,1,opt,name=info" json:"info,omitempty"`
	Kind             *socialaccountkindfilter.SocialAccountKindFilter `protobuf:"bytes,2,opt,name=kind" json:"kind,omitempty"`
	FollowersCount   *intfilter.IntFilter                             `protobuf:"bytes,3,opt,name=followersCount" json:"followersCount,omitempty"`
	FollowersSome    *SocialAccountFilter                             `protobuf:"bytes,4,opt,name=followersSome" json:"followersSome,omitempty"`
	FollowersEvery   *SocialAccountFilter                             `protobuf:"bytes,5,opt,name=followersEvery" json:"followersEvery,omitempty"`
	FollowersNone    *SocialAccountFilter                             `protobuf:"bytes,6,opt,name=followersNone" json:"followersNone,omitempty"`
	FollowingCount   *intfilter.IntFilter                             `protobuf:"bytes,7,opt,name=followingCount" json:"followingCount,omitempty"`
	FollowingSome    *SocialAccountFilter                             `protobuf:"bytes,8,opt,name=followingSome" json:"followingSome,omitempty"`
	FollowingEvery   *SocialAccountFilter                             `protobuf:"bytes,9,opt,name=followingEvery" json:"followingEvery,omitempty"`
	FollowingNone    *SocialAccountFilter                             `protobuf:"bytes,10,opt,name=followingNone" json:"followingNone,omitempty"`
	FriendsCount     *intfilter.IntFilter                             `protobuf:"bytes,11,opt,name=friendsCount" json:"friendsCount,omitempty"`
	FriendsSome      *SocialAccountFilter                             `protobuf:"bytes,12,opt,name=friendsSome" json:"friendsSome,omitempty"`
	FriendsEvery     *SocialAccountFilter                             `protobuf:"bytes,13,opt,name=friendsEvery" json:"friendsEvery,omitempty"`
	FriendsNone      *SocialAccountFilter                             `protobuf:"bytes,14,opt,name=friendsNone" json:"friendsNone,omitempty"`
	Feed             *FeedFilter                                      `protobuf:"bytes,15,opt,name=feed" json:"feed,omitempty"`
	PostsSome        *PostFilter                                      `protobuf:"bytes,16,opt,name=postsSome" json:"postsSome,omitempty"`
	PostsEvery       *PostFilter                                      `protobuf:"bytes,17,opt,name=postsEvery" json:"postsEvery,omitempty"`
	PostsNone        *PostFilter                                      `protobuf:"bytes,18,opt,name=postsNone" json:"postsNone,omitempty"`
	IdsSome          *idfilter.IdFilter                               `protobuf:"bytes,19,opt,name=idsSome" json:"idsSome,omitempty"`
	IdsEvery         *idfilter.IdFilter                               `protobuf:"bytes,20,opt,name=idsEvery" json:"idsEvery,omitempty"`
	IdsNone          *idfilter.IdFilter                               `protobuf:"bytes,21,opt,name=idsNone" json:"idsNone,omitempty"`
	Meta             *servicefilter.TypeMetaFilter                    `protobuf:"bytes,22,opt,name=meta" json:"meta,omitempty"`
	And              []*SocialAccountFilter                           `protobuf:"bytes,23,rep,name=and" json:"and,omitempty"`
	Or               []*SocialAccountFilter                           `protobuf:"bytes,24,rep,name=or" json:"or,omitempty"`
	Not              []*SocialAccountFilter                           `protobuf:"bytes,25,rep,name=not" json:"not,omitempty"`
	Hash             *string                                          `protobuf:"bytes,26,opt,name=hash" json:"hash,omitempty"`
	XXX_unrecognized []byte                                           `json:"-"`
}

func (m *SocialAccountFilter) Reset()                    { *m = SocialAccountFilter{} }
func (m *SocialAccountFilter) String() string            { return proto.CompactTextString(m) }
func (*SocialAccountFilter) ProtoMessage()               {}
func (*SocialAccountFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SocialAccountFilter) GetInfo() *infofilter.InfoFilter {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *SocialAccountFilter) GetKind() *socialaccountkindfilter.SocialAccountKindFilter {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *SocialAccountFilter) GetFollowersCount() *intfilter.IntFilter {
	if m != nil {
		return m.FollowersCount
	}
	return nil
}

func (m *SocialAccountFilter) GetFollowersSome() *SocialAccountFilter {
	if m != nil {
		return m.FollowersSome
	}
	return nil
}

func (m *SocialAccountFilter) GetFollowersEvery() *SocialAccountFilter {
	if m != nil {
		return m.FollowersEvery
	}
	return nil
}

func (m *SocialAccountFilter) GetFollowersNone() *SocialAccountFilter {
	if m != nil {
		return m.FollowersNone
	}
	return nil
}

func (m *SocialAccountFilter) GetFollowingCount() *intfilter.IntFilter {
	if m != nil {
		return m.FollowingCount
	}
	return nil
}

func (m *SocialAccountFilter) GetFollowingSome() *SocialAccountFilter {
	if m != nil {
		return m.FollowingSome
	}
	return nil
}

func (m *SocialAccountFilter) GetFollowingEvery() *SocialAccountFilter {
	if m != nil {
		return m.FollowingEvery
	}
	return nil
}

func (m *SocialAccountFilter) GetFollowingNone() *SocialAccountFilter {
	if m != nil {
		return m.FollowingNone
	}
	return nil
}

func (m *SocialAccountFilter) GetFriendsCount() *intfilter.IntFilter {
	if m != nil {
		return m.FriendsCount
	}
	return nil
}

func (m *SocialAccountFilter) GetFriendsSome() *SocialAccountFilter {
	if m != nil {
		return m.FriendsSome
	}
	return nil
}

func (m *SocialAccountFilter) GetFriendsEvery() *SocialAccountFilter {
	if m != nil {
		return m.FriendsEvery
	}
	return nil
}

func (m *SocialAccountFilter) GetFriendsNone() *SocialAccountFilter {
	if m != nil {
		return m.FriendsNone
	}
	return nil
}

func (m *SocialAccountFilter) GetFeed() *FeedFilter {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *SocialAccountFilter) GetPostsSome() *PostFilter {
	if m != nil {
		return m.PostsSome
	}
	return nil
}

func (m *SocialAccountFilter) GetPostsEvery() *PostFilter {
	if m != nil {
		return m.PostsEvery
	}
	return nil
}

func (m *SocialAccountFilter) GetPostsNone() *PostFilter {
	if m != nil {
		return m.PostsNone
	}
	return nil
}

func (m *SocialAccountFilter) GetIdsSome() *idfilter.IdFilter {
	if m != nil {
		return m.IdsSome
	}
	return nil
}

func (m *SocialAccountFilter) GetIdsEvery() *idfilter.IdFilter {
	if m != nil {
		return m.IdsEvery
	}
	return nil
}

func (m *SocialAccountFilter) GetIdsNone() *idfilter.IdFilter {
	if m != nil {
		return m.IdsNone
	}
	return nil
}

func (m *SocialAccountFilter) GetMeta() *servicefilter.TypeMetaFilter {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (m *SocialAccountFilter) GetAnd() []*SocialAccountFilter {
	if m != nil {
		return m.And
	}
	return nil
}

func (m *SocialAccountFilter) GetOr() []*SocialAccountFilter {
	if m != nil {
		return m.Or
	}
	return nil
}

func (m *SocialAccountFilter) GetNot() []*SocialAccountFilter {
	if m != nil {
		return m.Not
	}
	return nil
}

func (m *SocialAccountFilter) GetHash() string {
	if m != nil && m.Hash != nil {
		return *m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*FeedFilter)(nil), "feedfilter.FeedFilter")
	proto.RegisterType((*PostFilter)(nil), "feedfilter.PostFilter")
	proto.RegisterType((*SocialAccountFilter)(nil), "feedfilter.SocialAccountFilter")
}

func init() { proto.RegisterFile("feedfilter/feedfilter_.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1130 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x58, 0xdf, 0x6e, 0xdb, 0xb6,
	0x17, 0x46, 0x12, 0xe7, 0x8f, 0x8f, 0xed, 0x38, 0x65, 0xd2, 0x54, 0xf5, 0x2f, 0x4d, 0xf3, 0x0b,
	0x86, 0x22, 0x08, 0x32, 0xba, 0x09, 0x0a, 0xaf, 0x28, 0xba, 0x8b, 0x3a, 0x6d, 0x86, 0xa0, 0x6b,
	0x31, 0x24, 0xb9, 0xda, 0xcd, 0xa0, 0x49, 0xb4, 0x4d, 0xd4, 0x16, 0x0d, 0x89, 0xe9, 0xea, 0x27,
	0xdb, 0x3b, 0xec, 0x69, 0xf6, 0x08, 0x03, 0x8f, 0x48, 0x89, 0x92, 0x5d, 0x59, 0xca, 0x1d, 0x49,
	0x7d, 0x1f, 0xf9, 0xf1, 0xf0, 0xe3, 0x39, 0xb4, 0xe1, 0x60, 0xc0, 0x98, 0x3f, 0xe0, 0x63, 0xc9,
	0xc2, 0x6e, 0xda, 0xfc, 0x83, 0x4e, 0x43, 0x21, 0x05, 0x81, 0x74, 0xa8, 0xf3, 0x83, 0x6a, 0x7f,
	0xe1, 0x81, 0x8d, 0x4e, 0xbb, 0x9a, 0xd1, 0x39, 0xe0, 0xc1, 0x40, 0x68, 0x44, 0xda, 0x34, 0x5f,
	0x3b, 0x3c, 0x90, 0xc9, 0x47, 0x99, 0xfd, 0xe6, 0x70, 0x33, 0x33, 0xcf, 0xcd, 0x79, 0x1c, 0xb1,
	0xf0, 0x2b, 0xf7, 0x98, 0xfe, 0x9c, 0xe9, 0x25, 0xeb, 0x4a, 0xf6, 0xcd, 0x4c, 0x9d, 0x36, 0xcd,
	0xd7, 0xb3, 0x31, 0xff, 0xc2, 0x24, 0x0b, 0x27, 0x96, 0xfe, 0xf9, 0x21, 0x83, 0xbe, 0xf0, 0xc4,
	0x64, 0xc2, 0x02, 0x99, 0x23, 0x2c, 0x1c, 0x35, 0x1c, 0x1a, 0x8d, 0xdc, 0x30, 0xbf, 0xc4, 0x82,
	0x31, 0x83, 0x3f, 0x9c, 0x30, 0x9f, 0xbb, 0x1a, 0x67, 0xb5, 0xcd, 0xf7, 0x5e, 0x24, 0x3c, 0xee,
	0x8e, 0x5d, 0xcf, 0x13, 0xf7, 0x81, 0xb4, 0xe7, 0x5c, 0x3c, 0xae, 0x79, 0xc7, 0xff, 0xac, 0x03,
	0x5c, 0x31, 0xe6, 0x5f, 0xe1, 0x28, 0xb9, 0x80, 0x9a, 0xc2, 0x38, 0x2b, 0x47, 0x2b, 0x27, 0x8d,
	0x8b, 0x43, 0x9a, 0x3d, 0x34, 0xaa, 0x90, 0x1f, 0x79, 0xa0, 0xd1, 0x37, 0x88, 0x25, 0xa7, 0x50,
	0x53, 0x27, 0xe7, 0xac, 0x22, 0x67, 0x9f, 0xa6, 0xc7, 0x48, 0xaf, 0x83, 0x81, 0x30, 0x58, 0x35,
	0x4c, 0x7e, 0x86, 0xba, 0x17, 0x32, 0x57, 0x32, 0xbf, 0x3f, 0x73, 0xd6, 0x90, 0xf0, 0x9c, 0xa6,
	0xa6, 0xa1, 0xb7, 0xa8, 0xf6, 0x5d, 0xac, 0x56, 0x33, 0x53, 0x06, 0x39, 0x85, 0x75, 0xfc, 0xe2,
	0xd4, 0x90, 0xba, 0x47, 0x13, 0x57, 0xd0, 0xeb, 0x04, 0x1f, 0x43, 0xc8, 0x2b, 0xa8, 0x4f, 0x45,
	0x24, 0xa3, 0x5b, 0x31, 0x61, 0xce, 0xba, 0xd6, 0x66, 0x2d, 0xf5, 0x9b, 0x88, 0x92, 0x15, 0x12,
	0x20, 0xe9, 0x01, 0x60, 0xe7, 0xc3, 0x57, 0x16, 0xce, 0x9c, 0x8d, 0x42, 0x9a, 0x85, 0x4c, 0x56,
	0xfb, 0x2c, 0x02, 0xe6, 0x6c, 0x96, 0x58, 0x4d, 0x01, 0xc9, 0x19, 0x6c, 0x72, 0x3f, 0x56, 0xb8,
	0x85, 0x1c, 0x42, 0x8d, 0x99, 0xe9, 0xb5, 0x89, 0xb2, 0x81, 0x10, 0x0a, 0x5b, 0xdc, 0xd7, 0xca,
	0xea, 0xdf, 0x85, 0x27, 0x18, 0x3d, 0x3b, 0x2a, 0x82, 0xc2, 0xd9, 0x51, 0xcb, 0x39, 0xd4, 0x26,
	0x4c, 0xba, 0x4e, 0x03, 0xa1, 0xcf, 0x68, 0xe6, 0xda, 0xd0, 0xbb, 0xd9, 0x94, 0x7d, 0x62, 0xd2,
	0x35, 0xa7, 0xa9, 0xa0, 0xe4, 0x04, 0xd6, 0xdc, 0xc0, 0x77, 0x9a, 0x47, 0x6b, 0xf9, 0xed, 0xa6,
	0x96, 0xba, 0x51, 0x10, 0xf2, 0x02, 0x56, 0x45, 0xe8, 0xb4, 0x0a, 0x81, 0xab, 0x22, 0x54, 0x33,
	0x06, 0x42, 0x3a, 0xdb, 0xc5, 0x33, 0x06, 0x42, 0x12, 0x02, 0xb5, 0x91, 0x1b, 0x8d, 0x9c, 0xf6,
	0xd1, 0xca, 0x49, 0xfd, 0x06, 0xdb, 0xc7, 0x7f, 0xb7, 0x01, 0xd2, 0x40, 0x93, 0x9f, 0x60, 0xc3,
	0xbd, 0x97, 0x23, 0x11, 0x6a, 0x3b, 0x2f, 0x75, 0x9a, 0x86, 0x93, 0x97, 0xb0, 0xe9, 0x89, 0x40,
	0xb2, 0x40, 0x26, 0xa6, 0x4e, 0x73, 0x04, 0xbd, 0x63, 0xdf, 0x0c, 0xc1, 0xc0, 0x48, 0x1f, 0xb6,
	0x54, 0x7e, 0xb8, 0x63, 0xe1, 0x44, 0xdb, 0xfa, 0x05, 0x9d, 0x4f, 0x18, 0xf4, 0x57, 0x8d, 0xb1,
	0xee, 0x50, 0xc2, 0x23, 0xaf, 0x00, 0x54, 0x3b, 0xba, 0x5c, 0xea, 0x70, 0x0b, 0x47, 0xde, 0x41,
	0x43, 0xf5, 0xfc, 0xfe, 0xcc, 0x32, 0xfa, 0xd2, 0x9d, 0xda, 0x1c, 0x72, 0x09, 0x4d, 0xdd, 0xb5,
	0x5d, 0xbf, 0x74, 0x8e, 0x0c, 0xc9, 0xd2, 0x61, 0x5d, 0x81, 0xd2, 0x3a, 0xd0, 0x81, 0x9f, 0xa1,
	0xa1, 0x73, 0x26, 0xc6, 0x31, 0xbe, 0x11, 0x67, 0x74, 0x61, 0x1e, 0xa5, 0x97, 0x29, 0xd2, 0x8a,
	0xa6, 0x3d, 0x01, 0x79, 0x03, 0x2d, 0xdd, 0xd5, 0x31, 0xad, 0x17, 0xc4, 0x34, 0x0b, 0x25, 0xd7,
	0xd0, 0xd6, 0x03, 0x49, 0x68, 0xa1, 0xdc, 0x96, 0xf2, 0x3c, 0xf2, 0x11, 0x76, 0xac, 0xa1, 0x38,
	0xc4, 0x8d, 0x72, 0x73, 0xcd, 0x11, 0x73, 0xba, 0x30, 0xd4, 0xcd, 0xea, 0xba, 0x30, 0xdc, 0x6f,
	0xa0, 0x69, 0xf6, 0x8c, 0xfb, 0x6b, 0x15, 0x66, 0xad, 0x0c, 0x96, 0xbc, 0x4d, 0x43, 0x1b, 0x6f,
	0x68, 0xbb, 0x90, 0x9c, 0x05, 0xdb, 0x2b, 0xe3, 0x0e, 0xda, 0xe5, 0x56, 0x46, 0xd5, 0x3d, 0x00,
	0xdd, 0xe7, 0xc1, 0xd0, 0xd9, 0x29, 0x4e, 0xd0, 0x29, 0x92, 0x5c, 0x41, 0x1d, 0xcb, 0x2b, 0x5a,
	0xeb, 0x11, 0xd2, 0x4e, 0x16, 0x15, 0x61, 0x7a, 0x6b, 0x50, 0x96, 0xad, 0x52, 0x2a, 0xe9, 0x41,
	0x03, 0x3b, 0xda, 0x52, 0xa4, 0xc0, 0x52, 0x36, 0x90, 0x50, 0xd8, 0xc0, 0xae, 0xef, 0xec, 0x16,
	0x6a, 0xd6, 0x28, 0x75, 0x29, 0xe3, 0x96, 0x76, 0xdf, 0x5e, 0xc9, 0x4b, 0x69, 0x93, 0xc8, 0x07,
	0x68, 0x99, 0x7e, 0x7c, 0x4c, 0x8f, 0xcb, 0xcd, 0x92, 0x65, 0xd9, 0x5a, 0xf0, 0xbc, 0xf6, 0x2b,
	0x6a, 0xc1, 0x83, 0xeb, 0x43, 0xdb, 0x95, 0xd2, 0xf5, 0x46, 0xa9, 0xe3, 0x9e, 0xe0, 0x3c, 0x0e,
	0xb5, 0xde, 0x33, 0xf4, 0x93, 0x6a, 0x1b, 0xcb, 0xe6, 0x08, 0xe4, 0x3d, 0xec, 0x58, 0x43, 0xf1,
	0x96, 0x9c, 0x25, 0x93, 0xcc, 0x31, 0x72, 0x4a, 0x70, 0x47, 0x4f, 0x2b, 0x28, 0xc9, 0x57, 0xee,
	0x4e, 0xb5, 0xca, 0xfd, 0xbf, 0x6a, 0x95, 0xfb, 0xa0, 0x7c, 0xe5, 0x7e, 0x56, 0xb9, 0x72, 0x1f,
	0xce, 0xd7, 0x59, 0xcb, 0x8a, 0x56, 0xe5, 0x7e, 0x5e, 0x08, 0xb4, 0x2a, 0xf7, 0x51, 0xf1, 0x8c,
	0x76, 0xe5, 0xfe, 0xbf, 0x55, 0xb9, 0xff, 0x6d, 0xc0, 0xee, 0x02, 0x0b, 0x25, 0x6f, 0xcb, 0x95,
	0x12, 0x6f, 0xcb, 0xf7, 0xfa, 0xed, 0x1a, 0x97, 0xec, 0x97, 0xf4, 0x3b, 0x2f, 0xdf, 0xac, 0x55,
	0xe7, 0x5e, 0xb3, 0x6f, 0x61, 0x7b, 0x20, 0xc6, 0x63, 0xf1, 0x17, 0x0b, 0xf5, 0x15, 0x5f, 0x2b,
	0xb8, 0xe2, 0x39, 0xac, 0xba, 0x70, 0xc9, 0x08, 0x9a, 0xa3, 0x56, 0xf2, 0xc2, 0x65, 0x58, 0xe4,
	0x17, 0x4b, 0x44, 0xec, 0x9a, 0x92, 0x75, 0x3d, 0x47, 0xcb, 0xe8, 0x41, 0x3b, 0x6d, 0x54, 0xd5,
	0x83, 0x0e, 0x4b, 0x82, 0xc2, 0x83, 0x61, 0x1c, 0x94, 0xcd, 0xe5, 0x41, 0x31, 0xd8, 0x54, 0x04,
	0x0f, 0x86, 0xd6, 0x5b, 0xb7, 0xac, 0x08, 0xcd, 0x4a, 0x83, 0xc2, 0x83, 0xa1, 0xfd, 0x08, 0x2e,
	0x1b, 0x14, 0x43, 0xcb, 0xe8, 0xb1, 0x5e, 0xc7, 0x15, 0xf4, 0x60, 0x50, 0x5e, 0x43, 0x73, 0x10,
	0x72, 0x16, 0xf8, 0xda, 0x27, 0x8d, 0x82, 0x90, 0x64, 0x90, 0xea, 0xad, 0xa4, 0xfb, 0x18, 0x8e,
	0x92, 0x05, 0xdc, 0xe6, 0xa8, 0x94, 0xac, 0xbb, 0x71, 0x28, 0x5a, 0x25, 0x53, 0xb2, 0x4d, 0xb2,
	0x74, 0x60, 0x18, 0xb6, 0xab, 0xe9, 0xc0, 0x20, 0x9c, 0x42, 0x4d, 0xc1, 0x17, 0x95, 0x70, 0xeb,
	0xc5, 0x8e, 0x98, 0xec, 0x2f, 0xb2, 0x9d, 0x87, 0xfd, 0x22, 0x7b, 0xf4, 0xb0, 0x5f, 0x64, 0xe4,
	0x01, 0xbf, 0xc8, 0x76, 0xab, 0xe5, 0xf5, 0xbd, 0x6a, 0x79, 0xfd, 0x71, 0xf9, 0xbc, 0xbe, 0x5f,
	0x3e, 0xaf, 0x9f, 0xc7, 0x79, 0xfd, 0x09, 0x66, 0xe1, 0xa5, 0x27, 0x89, 0x09, 0xbe, 0x8b, 0x09,
	0xde, 0x29, 0xc7, 0x50, 0x99, 0xfe, 0x3c, 0xce, 0xf4, 0x4f, 0x4b, 0xae, 0x61, 0xa7, 0xfc, 0x4e,
	0x9a, 0xf2, 0xfb, 0xaf, 0x7f, 0xef, 0x0d, 0xb9, 0x1c, 0xdd, 0xff, 0xa9, 0x1e, 0xf8, 0xdd, 0x8b,
	0xf3, 0x48, 0x72, 0xd1, 0x1d, 0x8a, 0x1f, 0xa7, 0x63, 0x77, 0x36, 0x0c, 0xc5, 0x7d, 0xe0, 0x77,
	0x87, 0xe1, 0xd4, 0xeb, 0x46, 0xde, 0x88, 0x4d, 0x5c, 0xeb, 0xcf, 0xa6, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0xfc, 0x9d, 0x8c, 0x4a, 0x84, 0x12, 0x00, 0x00,
}