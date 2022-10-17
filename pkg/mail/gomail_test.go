package mail

import "testing"

func TestSendGoMail(t *testing.T) {
	type args struct {
		mailAddress []string
		subject     string
		body        string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "01",
			args: args{
				subject:     "您好，这是gomail-text测试邮件",
				mailAddress: []string{"2398027035@qq.com"},
				body: `
					您的服务存在异常，请查收！！！
				`},
			wantErr: false},
		{
			name: "02",
			args: args{
				subject:     "您好，这是gomail-html测试邮件",
				mailAddress: []string{"2398027035@qq.com"},
				body: `<html>
							<body>
								<h1>您的服务存在异常</h1>
							</body>
						</html>
				`},
			wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendGoMail(tt.args.mailAddress, tt.args.subject, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("SendGoMail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
