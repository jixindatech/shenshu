import store from '@/store'

export default {
  inserted(el, binding) {
    const { value } = binding
    const apis = store.getters && store.getters.api

    if (value && value instanceof Array && value.length > 0) {
      const hasPermission = value.every(val => apis.includes(val))
      if (!hasPermission) {
        el.parentNode && el.parentNode.removeChild(el)
      }
    } else {
      throw new Error('没有权限操作，请添加权限')
    }
  }
}
