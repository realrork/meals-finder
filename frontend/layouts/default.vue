<template>
  <div class="bg-base-200 px-[5%]">
    <div class="border-x-1 border-gray-700">
      <NavigationBar />
      <div class="bg-base-300 min-h-[79vh] p-2">
        <slot />
      </div>
      <FooterBar />
    </div>
  </div>
</template>

<script lang="ts" setup>

</script>

<style></style>