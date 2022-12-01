import request from '@/utils/request';
export function GetUserMessages(params) {
    return request({
        url: '/notify/message',
        method: 'GET',
        params: params
    });
}
export function UpdateUserMessage(params) {
    return request({
        url: '/notify/message/' + params.id,
        method: 'PATCH',
        data: params
    });
}
export function MessageNotify(params) {
    return request({
        url: '/notify/notify',
        method: 'POST',
        data: params
    });
}
export function DeleteMessage(id) {
    return request({
        url: '/notify/message/' + id,
        method: 'DELETE',
        data: {
            id
        }
    });
}
export function GetMessageTemplates(params) {
    return request({
        url: '/notify/message_template',
        method: 'GET',
        params: params
    });
}
export function CreateMessageTemplate(params) {
    return request({
        url: '/notify/message_template',
        method: 'POST',
        data: params
    });
}
export function UpdateMessageTemplate(params) {
    return request({
        url: '/notify/message_template/' + params.id,
        method: 'PATCH',
        data: params
    });
}
export function DeleteMessageTemplate(id) {
    return request({
        url: '/notify/message_template/' + id,
        method: 'DELETE',
        data: {
            id
        }
    });
}
//# sourceMappingURL=message.js.map