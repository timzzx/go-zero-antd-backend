declare namespace TAPI {

    type LoginParams = {
        name?: string;
        password?: string;
    };
    type LoginResult = {
        code?: number;
        msg?: string;
        token?: string;
    };

    type User = {
        id?: number;
        name?: string;
        role_id?: number;
        role_name?: string;
        ctime?: number;
        utime?: number;
    };

    type UserListParams = {
        name?: string;
    };
    type UserListResponse = {
        code?: number;
        msg?: string;
        data?: User[];
    }

    // 用户编辑
    type UserAddParams = {
        id?: number;
        name?: string;
        password?: string;
        role_id?: number;
    }
    type UserAddResponse = {
        code?: number;
        msg?: string;
    }
    // 用户删除
    type UserDelParams = {
        id?: number;
    }
    type UserDelResponse = {
        code?: number;
        msg?: string;
    }
    // 角色
    type Role = {
        id?: number;
        name?: string;
        type?: number;
        ctime?: number;
        utime?: number;
    };
    // 角色列表
    type RoleListResponse = {
        code?: number;
        msg?: string;
        data?: Role[];
    }
    // 角色编辑
    type RoleEditRequest = {
        id?: number;
        name?: string;
        type?: number;
    }
    type RoleEditResponse = {
        code?: number;
        msg?: string;
    }
    type RoleDelRequest = {
        id?: number;
    }
    type RoleDelResponse = {
        code?: number;
        msg?: string;
    }

    // 角色的资源列表
    type RolePermissionListParam = {
        role_id?: number;
    }
    type RolePermission = {
        id?: number;
        name?: string;
    }
    type RolePermissionListResponse = {
        code?: number;
        msg?: string;
        data?: RolePermission[];
        checkdata?: RolePermission[];
    }
    // 分配角色资源
    type rolePermissionEditParam = {
        role_id?: number;
        data?: string;
    }
    type rolePermissionEditResponse = {
        code?: number;
        msg?: string;
    }


    // 路由列表
    type Router = {
        method?: string;
        path?: string;
        name?: string;
    }
    type RouterListResponse = {
        code?: number;
        msg?: string;
        data?: Router[];
    }

    // 权限资源新增
    type RouterAddParam = {
        name?: string;
        path?: string;
    }
    type RouterAddResponse = {
        code?: number;
        msg?: string;
    }

}