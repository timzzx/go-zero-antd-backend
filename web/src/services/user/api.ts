// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** 登录接口 POST /api/login/account   表单提交*/
export async function login(body: USER.LoginParams, options?: { [key: string]: any }) {
    return request<USER.LoginResult>('/api/login', {
        method: 'POST',
        headers: {
        'Content-Type': 'application/x-www-form-urlencoded',
        },
        data: body,
        ...(options || {}),
    });
}

// 修改密码 json提交
export async function editPassword(body: USER.EditPasswordParam, options?: { [key: string]: any }) {
    return request<USER.EditPasswordResult>('/api/user/password',{
        method: 'POST',
        headers: {
        'Content-Type': 'application/json',
        },
        data: body, 
        ...(options || {}),
    });
}