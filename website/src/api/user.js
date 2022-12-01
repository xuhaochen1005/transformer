import request from '@/utils/request';
export function login(userInfo) {
    return request({
        url: '/user/token',
        method: 'POST',
        data: userInfo
    });
}
export function CreateUser(user) {
    return request({
        url: '/user/user',
        method: 'POST',
        data: user
    });
}
export function UpdateUser(user) {
    return request({
        url: '/user/' + user.id,
        method: 'PATCH',
        data: user
    });
}
export function GetUsers(query) {
    return request({
        url: '/user/',
        method: 'GET',
        params: query
    });
}
export function DeleteUser(id) {
    return request({
        url: '/user/' + id,
        method: 'DELETE'
    });
}
export function GetUserBasic() {
    return request({
        url: '/user/basic',
        method: 'GET'
        //  params: query
    });
}
//# sourceMappingURL=user.js.map