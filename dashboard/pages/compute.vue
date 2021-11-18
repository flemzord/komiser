<template>
  <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200">
    <div class="container mx-auto px-6 py-8">
      <h3 class="text-gray-700 text-3xl font-medium">Compute</h3>

      <div class="mt-4">
        <div class="flex flex-wrap -mx-6">
          <div class="w-full px-6 sm:w-1/2 xl:w-1/3">
            <div
              class="flex items-center px-5 py-6 shadow-sm rounded-md bg-white"
            >
              <div class="p-3 rounded-full bg-indigo-600 bg-opacity-75">
                <font-awesome-icon :icon="['fas', 'dollar-sign']" />
              </div>

              <div class="mx-5">
                <h4 class="text-2xl font-semibold text-gray-700">
                  <money-format
                    :value="currentCost"
                    locale="en"
                    currency-code="USD"
                    subunits-value="true"
                    hide-subunits="true"
                  >
                  </money-format>
                </h4>
                <div class="text-gray-500">Current month-to-date balance</div>
              </div>
            </div>
          </div>

          <div class="w-full px-6 sm:w-1/2 xl:w-1/3">
            <div
              class="flex items-center px-5 py-6 shadow-sm rounded-md bg-white"
            >
              <div class="p-3 rounded-full bg-indigo-600 bg-opacity-75">
                <font-awesome-icon :icon="['fas', 'ticket-alt']" />
              </div>

              <div class="mx-8">
                <h4 class="text-2xl font-semibold text-gray-700">
                  {{ openedTicket }} Open Tickets
                </h4>
                <p class="ext-2xl font-semibold text-gray-700">
                  {{ resolvedTicket }} Resolved Tickets
                </p>
                <div class="text-gray-500">Support</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-8"></div>

      <div class="flex flex-col mt-8">
        <div
          class="-my-2 py-2 overflow-x-auto sm:-mx-6 sm:px-6 lg:-mx-8 lg:px-8"
        >
          <div
            class="
              align-middle
              inline-block
              min-w-full
              shadow
              overflow-hidden
              sm:rounded-lg
              border-b border-gray-200
            "
          >
            <table class="min-w-full">
              <thead>
                <tr>
                  <th
                    class="
                      px-6
                      py-3
                      border-b border-gray-200
                      bg-gray-50
                      text-left text-xs
                      leading-4
                      font-medium
                      text-gray-500
                      uppercase
                      tracking-wider
                    "
                  >
                    Name
                  </th>
                  <th
                    class="
                      px-6
                      py-3
                      border-b border-gray-200
                      bg-gray-50
                      text-left text-xs
                      leading-4
                      font-medium
                      text-gray-500
                      uppercase
                      tracking-wider
                    "
                  >
                    Amount
                  </th>
                </tr>
              </thead>

              <tbody class="bg-white">
                <tr v-for="item in groupsCostHistory" :key="item.name">
                  <td
                    class="
                      px-6
                      py-4
                      whitespace-no-wrap
                      border-b border-gray-200
                    "
                  >
                    <div class="flex items-center">
                      <!--                      <div class="flex-shrink-0 h-10 w-10">-->
                      <!--                        <img class="h-10 w-10 rounded-full"-->
                      <!--                             src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&amp;ixid=eyJhcHBfaWQiOjEyMDd9&amp;auto=format&amp;fit=facearea&amp;facepad=2&amp;w=256&amp;h=256&amp;q=80"-->
                      <!--                             alt="">-->
                      <!--                      </div>-->

                      <div class="ml-4">
                        <div
                          class="text-sm leading-5 font-medium text-gray-900"
                        >
                          {{ item.name }}
                        </div>
                      </div>
                    </div>
                  </td>

                  <td
                    class="
                      px-6
                      py-4
                      whitespace-no-wrap
                      border-b border-gray-200
                    "
                  >
                    <div class="text-sm leading-5 text-gray-900">
                      <money-format
                        :value="item.cost"
                        locale="en"
                        currency-code="USD"
                        subunits-value="true"
                        hide-subunits="true"
                      />
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script>
import MoneyFormat from 'vue-money-format'

export default {
  components: {
    'money-format': MoneyFormat,
  },
  data() {
    return {
      openedTicket: '0',
      resolvedTicket: '0',
      currentCost: 0,
      costHistory: [{ groups: [] }],
    }
  },
  async created() {
    this.openedTicket = await this.$axios
      .$get('/aws/support/open', {
        progress: false,
      })
      .then((res) => {
        if (res.error) {
          return 0
        }
        return res
      })
    this.resolvedTicket = await this.$axios
      .$get('/aws/support/history', {
        progress: false,
      })
      .then((res) => {
        if (res.error) {
          return 0
        }
        return res
      })
    this.costHistory = await this.$axios.$get('/aws/cost/history', {
      progress: false,
    })
    this.currentCost = await this.$axios.$get('/aws/cost/current', {
      progress: false,
    })
  },
  computed: {
    groupsCostHistory: function () {
      this.mostUsedServices = []
      this.costHistory[this.costHistory.length - 1].groups
        .slice(0, 4)
        .forEach((service) => {
          this.mostUsedServices.push({
            name: service.key,
            cost: service.amount,
          })
        })
      return this.mostUsedServices
    },
  },
}
</script>
