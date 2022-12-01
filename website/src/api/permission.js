import request from '@/utils/request';
export function getRoles() {
    return request({
        url: '/user/roles',
        method: 'GET'
    });
}
export function getRoleList(params) {
    return request({
        url: '/permission/roles',
        method: 'GET',
        params: params
    });
}
export function getPermissions() {
    return request({
        url: '/user/permissions',
        method: 'GET'
    });
}
export function getRolePermissions(params) {
    return request({
        url: '/permission/role_permission',
        method: 'GET',
        params: params
    });
}
export function getUserRole(params) {
    return request({
        url: '/permission/user_role',
        method: 'GET',
        params: params
    });
}
export function updateUserRole(params) {
    return request({
        url: '/permission/user_role/' + params.id,
        method: 'PATCH',
        data: params
    });
}
export function createUserRole(params) {
    return request({
        url: '/permission/user_role',
        method: 'POST',
        data: params
    });
}
export function deleteUserRole(id) {
    return request({
        url: '/permission/user_role/' + id,
        method: 'DELETE'
    });
}
export function createRolePermissions(params) {
    return request({
        url: '/permission/role_permission',
        method: 'POST',
        data: {
            role_permissions: params
        }
    });
}
export function deleteRolePermission(params) {
    return request({
        url: '/permission/role_permission/' + [params].flatMap(i => i).join(','),
        method: 'DELETE'
    });
}
//# sourceMappingURL=permission.js.map