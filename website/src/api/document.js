import request from '@/utils/request';
export function GetDocuments(query) {
    return request({
        url: '/documents/',
        method: 'GET',
        params: query
    });
}
//# sourceMappingURL=document.js.map