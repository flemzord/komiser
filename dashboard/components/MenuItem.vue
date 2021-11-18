<template>
  <nuxt-link :to=url class="
        flex
        items-center
        mt-4
        py-2
        px-6
        text-gray-500
        hover:bg-gray-700 hover:bg-opacity-25 hover:text-gray-100
    " >
    <font-awesome-icon :icon=fa />
    <span class="mx-3">{{ name }}</span>
  </nuxt-link>
</template>

<script>
export default {
  props: ['name', "url", "fa"]
}
</script>
