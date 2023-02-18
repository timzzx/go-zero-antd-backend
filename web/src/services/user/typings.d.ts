declare namespace USER {
    
    type LoginParams = {
        name?: string;
        password?: string;
    };
    type LoginResult = {
        code?: number;
        msg?: string;
        token?: string;
    };

    // 修改密码参数
    type EditPasswordParam = {
        password: string;
    } 
    // 修改密码结果
    type EditPasswordResult = {
        code?: number;
        msg?: string;
    }

}