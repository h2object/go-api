package api

import (
	"github.com/h2object/rpc"
	"path"
	"time"
)

type Application struct{
	AppID 		string 	`json:"appid"`
	Secret 		string  `json:"secret"`
	Index 		string  `json:"index"`
	Admin 		bool    `json:"admin"`
	Cache 		bool    `json:"cache"`
	DefaultLimitSize int64 `json:"default_limit_size"`
}

func (h2o *H2Object) GetApplication(l Logger, auth Auth, app *Application) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("application", ".conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, app); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetApplication(l Logger, auth Auth, app *Application) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("application", ".conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, app, nil); err != nil {
		return err
	}
	return nil
}

type Service struct{
	Object 	bool 	`json:"object"`
	File 	bool 	`json:"file"`
	Auth 	bool 	`json:"auth"`
	Cache 	bool 	`json:"cache"`
	Token 	bool 	`json:"token"`
	System 	bool 	`json:"system"`
}

func (h2o *H2Object) GetService(l Logger, auth Auth, service *Service) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("service", ".conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, service); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetService(l Logger, auth Auth, service *Service) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("service", ".conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, service, nil); err != nil {
		return err
	}
	return nil
}

type EventConf struct{
	Currency int64 			`json:"currency"`			
}

func (h2o *H2Object) GetEvent(l Logger, auth Auth, event *EventConf) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("event", ".conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, event); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetEvent(l Logger, auth Auth, event *EventConf) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("event", ".conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, event, nil); err != nil {
		return err
	}
	return nil
}

type SMTP struct{
	Host string		`structs:"host"`
	Port int64		`structs:"port"`
	User string		`structs:"user"`
	Password string `structs:"password"`
	Sender string 	`structs:"sender"`
	ReplyTo string 	`structs:"replyto"`
}


func (h2o *H2Object) GetSMTP(l Logger, auth Auth, smtp *SMTP) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("smtp", ".conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, smtp); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetSMTP(l Logger, auth Auth, smtp *SMTP) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("smtp", ".conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, smtp, nil); err != nil {
		return err
	}
	return nil
}

type Session struct{
	Expire 	string 		`json:"duration"`
}

type SessionInternal struct{
	Duration float64 `json:"duration"`
}

func (h2o *H2Object) GetSession(l Logger, auth Auth, ssn *Session) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("session", ".conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	var ssni SessionInternal
	if err := h2o.conn.Get(l, URL, &ssni); err != nil {
		return err
	}

	ssn.Expire = time.Duration(ssni.Duration).String()
	return nil
}

func (h2o *H2Object) SetSession(l Logger, auth Auth, ssn *Session) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("session", ".conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	var ssni SessionInternal 
	duration, err := time.ParseDuration(ssn.Expire)
	if err != nil {
		return err
	}
	ssni.Duration = float64(duration.Nanoseconds())
	if err := h2o.conn.PutJson(l, URL, &ssni, nil); err != nil {
		return err
	}
	return nil
}

type Subject struct{
	Regist  string  		`json:"regist"`
	Active  string  		`json:"active"`
	Forget  string  		`json:"forget"`			
}

func (h2o *H2Object) GetSubject(l Logger, auth Auth, subject *Subject) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("subject", ".conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, subject); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetSubject(l Logger, auth Auth, subject *Subject) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("subject", ".conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, subject, nil); err != nil {
		return err
	}
	return nil
}

type TemplateConf struct{
	Markdown  string  		`json:"markdown"`
	Regist  string  		`json:"regist"`
	Active  string  		`json:"active"`
	Forget  string  		`json:"forget"`			
}

func (h2o *H2Object) GetTemplate(l Logger, auth Auth, tpl *TemplateConf) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("template", ".conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, tpl); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetTemplate(l Logger, auth Auth, tpl *TemplateConf) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("template", ".conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, tpl, nil); err != nil {
		return err
	}
	return nil
}

type AuthMobile struct{
	Enable bool 			`json:"enable"`
}

func (h2o *H2Object) GetAuthMobile(l Logger, auth Auth, mobile *AuthMobile) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "mobile.conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, mobile); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetAuthMobile(l Logger, auth Auth, mobile *AuthMobile) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "mobile.conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, mobile, nil); err != nil {
		return err
	}
	return nil
}

type AuthEmail struct{
	Enable bool 			`json:"enable"`
	Active bool 			`json:"active"`
}

func (h2o *H2Object) GetAuthEmail(l Logger, auth Auth, email *AuthEmail) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "email.conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, email); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetAuthEmail(l Logger, auth Auth, email *AuthEmail) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "email.conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, email, nil); err != nil {
		return err
	}
	return nil
}

type AuthQQ struct{
	Enable bool 		`json:"enable"`
	AppID  string 		`json:"appid"`
	Secret string		`json:"secret"`	
}

func (h2o *H2Object) GetAuthQQ(l Logger, auth Auth, qq *AuthQQ) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "qq.conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, qq); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetAuthQQ(l Logger, auth Auth, qq *AuthQQ) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "qq.conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, qq, nil); err != nil {
		return err
	}
	return nil
}

type AuthWeibo struct{
	Enable bool 		`json:"enable"`
	AppID  string 		`json:"appid"`
	Secret string		`json:"secret"`	
}

func (h2o *H2Object) GetAuthWeibo(l Logger, auth Auth, wb *AuthWeibo) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "weibo.conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, wb); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetAuthWeibo(l Logger, auth Auth, wb *AuthWeibo) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "weibo.conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, wb, nil); err != nil {
		return err
	}
	return nil
}

type AuthWechat struct{
	Enable bool 		`json:"enable"`
	AppID  string 		`json:"appid"`
	Secret string		`json:"secret"`	
}

func (h2o *H2Object) GetAuthWechat(l Logger, auth Auth, wechat *AuthWechat) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "wechat.conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, wechat); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetAuthWechat(l Logger, auth Auth, wechat *AuthWechat) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("auth", "wechat.conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, wechat, nil); err != nil {
		return err
	}
	return nil
}

type ThirdQiniu struct{
	Enable bool 		`json:"enable"`
	AppID  string		`json:"appid"`
	Secret string		`json:"secret"`
	Domain string 		`json:"domain"`
	Bucket string		`json:"bucket"`
}

func (h2o *H2Object) GetThirdQiniu(l Logger, auth Auth, qiniu *ThirdQiniu) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("third", "qiniu.conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, qiniu); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetThirdQiniu(l Logger, auth Auth, qiniu *ThirdQiniu) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("third", "qiniu.conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, qiniu, nil); err != nil {
		return err
	}
	return nil
}

type FileLogConf struct {
	Enable bool 		`json:"enable"`
	FileName string 	`json:"filename"`
	Level    string	 	`json:"level"`
	Rotate   bool 		`json:"rotate"`
	RotateDaily bool 	`json:"rotate_daily"`
	RotateMaxLine int64 	`json:"rotate_max_line"`
	RotateMaxSize int64 	`json:"rotate_max_size"`
}

func (h2o *H2Object) GetFileLogConf(l Logger, auth Auth, file *FileLogConf) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("log", ".conf"), nil)

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.Get(l, URL, file); err != nil {
		return err
	}
	return nil
}

func (h2o *H2Object) SetFileLogConf(l Logger, auth Auth, file *FileLogConf) error {
	URL := rpc.BuildHttpURL(h2o.addr, path.Join("log", ".conf"), nil)
	if l != nil {
		l.Debug("URL: %s", URL)
	}

	h2o.Lock()
	defer h2o.Unlock()

	h2o.conn.Prepare(auth)
	defer h2o.conn.Prepare(nil)

	if err := h2o.conn.PutJson(l, URL, file, nil); err != nil {
		return err
	}
	return nil
}



