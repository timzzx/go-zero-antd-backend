// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

/** 登录接口 POST /api/login/account */
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