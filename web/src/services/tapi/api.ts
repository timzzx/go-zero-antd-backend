// @ts-ignore
/* eslint-disable */
import { request } from '@umijs/max';

// 修改密码 json提交
export async function userList(body: TAPI.UserListParams | undefined = undefined, options?: { [key: string]: any }) {
    return request<TAPI.UserListResponse>('/api/user/list', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        data: body,
        ...(options || {}),
    });
}

// 角色列表
export async function roleList() {
    return request<TAPI.RoleListResponse>("/api/role/list", {
        method: "GET",
    });
}

// 角色编辑
export async function roleEdit(params: TAPI.RoleEditRequest) {
    return request<TAPI.RoleEditResponse>("/api/role/edit", {
        method: "POST",
        params: params,
    })
}

// 角色删除
export async function roleDel(params: TAPI.RoleDelRequest) {
    return request<TAPI.RoleDelResponse>("/api/role/delete", {
        method: "POST",
        params: params,
    })
}

// 角色权限列表
export async function rolePermissionList(params: TAPI.RolePermissionListParam) {
    return request<TAPI.RolePermissionListResponse>("/api/role/permission/resource/list", {
        method: "POST",
        params: params,
    })
}

export async function rolePermissionEdit(params: TAPI.rolePermissionEditParam) {
    return request<TAPI.rolePermissionEditResponse>("/api/role/permission/resource/edit", {
        method: "POST",
        params: params,
    })
}

// 用户编辑
export async function userAdd(params: TAPI.UserAddParams) {
    return request<TAPI.UserAddResponse>("/api/user/add", {
        method: "POST",
        params: params,
    });
}

// 用户删除
export async function userDel(params: TAPI.UserDelParams) {
    return request<TAPI.UserDelResponse>("/api/user/del", {
        method: "POST",
        data: params,
    });
}

// 路由列表
export async function routerList() {
    return request<TAPI.RouterListResponse>("/api/router/list", {
        method: "GET",
    })
}

// 新增权限
export async function routerAdd(data: TAPI.RouterAddParam) {
    return request<TAPI.RouterAddResponse>("/api/router/add", {
        method: "POST",
        data: data,
    })
}