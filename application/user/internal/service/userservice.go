package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/transport"
	"user/internal/biz"

	pb "user/api/user/v1"
)

type UserServiceService struct {
	pb.UnimplementedUserServiceServer
	uc *biz.UserUsecase
}

func NewUserServiceService(uc *biz.UserUsecase) *UserServiceService {
	return &UserServiceService{
		uc: uc,
	}
}

func (s *UserServiceService) Signin(ctx context.Context, req *pb.SigninRequest) (*pb.SigninReply, error) {
	result, err := s.uc.Signin(ctx, &biz.SigninRequest{
		State: req.State,
		Code:  req.Code,
	})
	fmt.Printf("result:%+v", result)
	if err != nil {
		return nil, err
	}
	return &pb.SigninReply{
		State: result.State,
		Data:  result.Data,
	}, nil
}

func (s *UserServiceService) GetUserInfo(ctx context.Context, req *pb.GetUserInfoRequest) (*pb.GetUserInfoResponse, error) {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		fmt.Println("获取header失败")
		return nil, errors.New("获取header失败")
	}
	header := tr.RequestHeader()
	authorization := header.Get("Authorization")

	result, err := s.uc.GetUserInfo(ctx, &biz.GetUserInfoRequest{
		Authorization: authorization,
	})
	if err != nil {
		return nil, err
	}
	return &pb.GetUserInfoResponse{
		State: result.State,
		Data: &pb.Data{
			Owner: result.Data.Owner,
			Name:  result.Data.Name,
			// CreatedTime:         result.Data.CreatedTime,
			// UpdatedTime:         result.Data.UpdatedTime,
			// DeletedTime:         result.Data.DeletedTime,
			Id:              result.Data.Id,
			ExternalId:      result.Data.ExternalId,
			Type:            result.Data.Type,
			Password:        result.Data.Password,
			PasswordSalt:    result.Data.PasswordSalt,
			PasswordType:    result.Data.PasswordType,
			DisplayName:     result.Data.DisplayName,
			FirstName:       result.Data.FirstName,
			LastName:        result.Data.LastName,
			Avatar:          result.Data.Avatar,
			AvatarType:      result.Data.AvatarType,
			PermanentAvatar: result.Data.PermanentAvatar,
			Email:           result.Data.Email,
			EmailVerified:   result.Data.EmailVerified,
			Phone:           result.Data.Phone,
			CountryCode:     result.Data.CountryCode,
			Region:          result.Data.Region,
			Location:        result.Data.Location,
			Address:         result.Data.Address,
			Affiliation:     result.Data.Affiliation,
			Title:           result.Data.Title,
			IdCardType:      result.Data.IdCardType,
			IdCard:          result.Data.IdCard,
			Homepage:        result.Data.Homepage,
			Bio:             result.Data.Bio,
			Tag:             result.Data.Tag,
			Language:        result.Data.Language,
			Gender:          result.Data.Gender,
			Birthday:        result.Data.Birthday,
			Education:       result.Data.Education,
			// Score:               result.Data.Score,
			// Karma:               result.Data.Karma,
			// Ranking:             result.Data.Ranking,
			// Balance:             result.Data.Balance,
			// Currency:            result.Data.Currency,
			IsDefaultAvatar:   result.Data.IsDefaultAvatar,
			IsOnline:          result.Data.IsOnline,
			IsAdmin:           result.Data.IsAdmin,
			IsForbidden:       result.Data.IsForbidden,
			IsDeleted:         result.Data.IsDeleted,
			SignupApplication: result.Data.SignupApplication,
			Hash:              result.Data.Hash,
			PreHash:           result.Data.PreHash,
			AccessKey:         result.Data.AccessKey,
			AccessSecret:      result.Data.AccessSecret,
			// AccessToken:         result.Data.AccessToken,
			CreatedIp:      result.Data.CreatedIp,
			LastSigninTime: result.Data.LastSigninTime,
			LastSigninIp:   result.Data.LastSigninIp,
			// Github:              result.Data.Github,
			Google: result.Data.Google,
			// Qq:                  result.Data.Qq,
			// Wechat:              result.Data.Wechat,
			Facebook: result.Data.Facebook,
			// Dingtalk:            result.Data.Dingtalk,
			Weibo: result.Data.Weibo,
			Gitee: result.Data.Gitee,
			// Linkedin:            result.Data.Linkedin,
			Wecom:    result.Data.Wecom,
			Lark:     result.Data.Lark,
			Gitlab:   result.Data.Gitlab,
			Adfs:     result.Data.Adfs,
			Baidu:    result.Data.Baidu,
			Alipay:   result.Data.Alipay,
			Casdoor:  result.Data.Casdoor,
			Infoflow: result.Data.Infoflow,
			Apple:    result.Data.Apple,
			// Azuread:             result.Data.Azuread,
			// Azureadb2C:          result.Data.Azureadb2C,
			Slack:    result.Data.Slack,
			Steam:    result.Data.Steam,
			Bilibili: result.Data.Bilibili,
			Okta:     result.Data.Okta,
			Douyin:   result.Data.Douyin,
			Line:     result.Data.Line,
			Amazon:   result.Data.Amazon,
			Auth0:    result.Data.Auth0,
			// Battlenet:           result.Data.Battlenet,
			Bitbucket: result.Data.Bitbucket,
			Box:       result.Data.Box,
			// Cloudfoundry:        result.Data.Cloudfoundry,
			Dailymotion: result.Data.Dailymotion,
			Deezer:      result.Data.Deezer,
			// Digitalocean:        result.Data.Digitalocean,
			Discord: result.Data.Discord,
			Dropbox: result.Data.Dropbox,
			// Eveonline:           result.Data.Eveonline,
			// Fitbit: result.Data.Fitbit,
			Gitea:  result.Data.Gitea,
			Heroku: result.Data.Heroku,
			// Influxcloud:         result.Data.Influxcloud,
			// Instagram: result.Data.Instagram,
			Intercom: result.Data.Intercom,
			Kakao:    result.Data.Kakao,
			Lastfm:   result.Data.Lastfm,
			Mailru:   result.Data.Mailru,
			Meetup:   result.Data.Meetup,
			// Microsoftonline:     result.Data.Microsoftonline,
			// Naver:               result.Data.Naver,
			Nextcloud: result.Data.Nextcloud,
			// Onedrive:            result.Data.Onedrive,
			Oura:    result.Data.Oura,
			Patreon: result.Data.Patreon,
			Paypal:  result.Data.Paypal,
			// Salesforce:          result.Data.Salesforce,
			Shopify:    result.Data.Shopify,
			Soundcloud: result.Data.Soundcloud,
			Spotify:    result.Data.Spotify,
			Strava:     result.Data.Strava,
			Stripe:     result.Data.Stripe,
			// Tiktok:              result.Data.Tiktok,
			Tumblr:   result.Data.Tumblr,
			Twitch:   result.Data.Twitch,
			Twitter:  result.Data.Twitter,
			Typetalk: result.Data.Typetalk,
			Uber:     result.Data.Uber,
			// Vk:                  result.Data.Vk,
			Wepay:  result.Data.Wepay,
			Xero:   result.Data.Xero,
			Yahoo:  result.Data.Yahoo,
			Yammer: result.Data.Yammer,
			Yandex: result.Data.Yandex,
			Zoom:   result.Data.Zoom,
			// Metamask:            result.Data.Metamask,
			Web3Onboard: result.Data.Web3Onboard,
			Custom:      result.Data.Custom,
			// WebauthnCredentials: result.Data.WebauthnCredentials,
			PreferredMfaType: result.Data.PreferredMfaType,
			// RecoveryCodes:       result.Data.RecoveryCodes,
			TotpSecret:      result.Data.TotpSecret,
			MfaPhoneEnabled: result.Data.MfaPhoneEnabled,
			MfaEmailEnabled: result.Data.MfaEmailEnabled,
			// Invitation:          result.Data.Invitation,
			// InvitationCode:      result.Data.InvitationCode,
			// FaceIds:             result.Data.FaceIds,
			Ldap: result.Data.Ldap,
			// Properties:          result.Data.Properties,
			// Roles:               result.Data.Roles,
			// Permissions:         result.Data.Permissions,
			// Groups:              result.Data.Groups,
			// SigninWrongTimes:    result.Data.SigninWrongTimes,
			LastSigninWrongTime: result.Data.LastSigninWrongTime,
			// ManagedAccounts:     result.Data.ManagedAccounts,
		},
	}, nil
}
