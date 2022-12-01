export var DeviceType;
(function (DeviceType) {
  DeviceType[DeviceType['Mobile'] = 0] = 'Mobile'
  DeviceType[DeviceType['Desktop'] = 1] = 'Desktop'
})(DeviceType || (DeviceType = {}))
export const appState = {
  device: DeviceType.Desktop,
  sidebar: {
    opened: localStorage.getItem('sidebar') !== 'closed',
    withoutAnimation: false
  }
}
function getToken() {
  const tokenStr = localStorage.getItem('token@' + document.domain)
  if (tokenStr) {
    const tokenObj = JSON.parse(tokenStr)
    if (new Date(tokenObj.expires).getTime() < new Date().getTime()) {
      localStorage.removeItem('token@' + document.domain)
      return ''
    } else {
      return tokenObj.token
    }
  } else {
    return ''
  }
}
export const userState = {
  id: -1,
  token: getToken(),
  name: '',
  avatar: '',
  introduction: '',
  telephone: '',
  email: ''
}
export const tagViewState = {
  visitedViews: [],
  cachedViews: []
}
// # sourceMappingURL=state.js.map
