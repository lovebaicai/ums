<template>
  <div class="dashboard-editor-container">
    <panel-group />
  </div>
</template>
<script>
import PanelGroup from './components/PanelGroup'

export default {
  name: 'Dashboard',
  components: {
    PanelGroup
  }
}
</script>

<style lang="scss" scoped>
  .dashboard-editor-container {
    padding: 32px;
    position: relative;
}

@media (max-width:1024px) {
  .chart-wrapper {
    padding: 8px;
  }
}
</style>
