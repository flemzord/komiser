import { mount } from '@vue/test-utils'
import Menu from '@/components/Menu.vue'

describe('Menu', () => {
  test('is a Vue instance', () => {
    const wrapper = mount(Menu)
    expect(wrapper.vm).toBeTruthy()
  })
})
