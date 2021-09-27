import { mount } from '@vue/test-utils'
import LogoText from '@/components/LogoText.vue'

describe('LogoText', () => {
  test('is a Vue instance', () => {
    const wrapper = mount(LogoText)
    expect(wrapper.vm).toBeTruthy()
  })
})
