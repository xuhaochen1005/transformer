import request from '@/utils/request';
export function GetStdLtiLibraries(query) {
    return request({
        url: '/std/lti',
        method: 'GET',
        params: query
    });
}
export function CreateStdLtiLibraries(lti) {
    return request({
        url: '/std/lti',
        method: 'POST',
        data: lti
    });
}
export function UpdateStdLtiLibraries(lti) {
    return request({
        url: '/std/lti/' + lti.id,
        method: 'PATCH',
        data: lti
    });
}
export function DeleteStdLtiLibraries(id) {
    return request({
        url: '/std/lti/' + id,
        method: 'DELETE'
    });
}
//# sourceMappingURL=lti.js.map