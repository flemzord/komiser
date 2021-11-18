<template>
  <nav class="mt-10">
    <MenuItem name="Dashboard" url="/" fa="house-damage" />
    <MenuItem name="Compute" url="/compute" fa="lock" />
    <MenuItem name="Storage" url="/storage" fa="warehouse" />
    <MenuItem name="Network" url="/network" fa="network-wired" />
    <MenuItem name="Security" url="/security" fa="shield-alt" />
    <MenuItem name="Data & AI" url="/data" fa="database" />
  </nav>
</template>
<script>
import MenuItem from "./MenuItem";
export default {
  components: {MenuItem}
}
</script>
