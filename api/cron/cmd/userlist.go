package cmd

import (
	"context"
	"fmt"
)

func userlist() {
	u := svcCtx.BkModel.User
	d, err := u.WithContext(context.Background()).Debug().First()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(d)
}
