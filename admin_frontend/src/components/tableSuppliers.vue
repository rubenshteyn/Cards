<template>
  <div class="d-flex flex-column flex-column-fluid container-fluid">
    <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
      <!--begin::Page title-->
      <div class="page-title d-flex flex-column me-3">
        <!--begin::Title-->
        <h1 class="d-flex text-dark fw-bold my-1 fs-3">Supplier management</h1>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">Home</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">Supplier management</li>
          <!--end::Item-->
        </ul>
        <!--end::Breadcrumb-->
      </div>
      <!--end::Page title-->
      <!--begin::Actions-->
<!--      <div class="actions-group-buttons d-flex align-items-center py-2 py-md-1 position-relative">-->
<!--        <button class="btn btn-dark fw-bold" data-bs-toggle="modal"-->
<!--                data-bs-target="#kt_modal_create_app" id="kt_toolbar_primary_button">Add supplier-->
<!--        </button>-->
<!--      </div>-->
      <!--end::Actions-->
    </div>
    <div class="card mb-5 mb-xl-8">
      <!--begin::Header-->
      <div class="card-header border-0 pt-5">
        <h3 class="card-title align-items-start flex-column">
          <span class="card-label fw-bold fs-3 mb-1">Suppliers</span>
          <span class="text-muted mt-1 fw-semibold fs-7">Amount suppliers: {{this.suppliers.length}}</span>
        </h3>
      </div>
      <!--end::Header-->
      <!--begin::Body-->
      <div class="card-body py-3">
        <!--begin::Table container-->
        <div class="table-responsive">
          <!--begin::Table-->
          <table v-if="getSuppliers()" class="table table-row-bordered table-row-gray-100 align-middle gs-0 gy-3">
            <!--begin::Table head-->
            <thead class="rounded-left">
            <tr class="fw-bold text-muted bg-light">
              <th class="min-w-150px">Supplier</th>
              <th @click="filterUserDate('CreatedAt')" class="min-w-150px cursor-pointer text-hover-primary">
                Registration
              </th>
              <th @click="filterUserDate('CreatedAt')" class="min-w-150px cursor-pointer text-hover-primary">
                Updates
              </th>
              <th class="min-w-120px">Status</th>
              <th class="min-w-100px text-end">Actions</th>
            </tr>
            </thead>
            <!--end::Table head-->
            <!--begin::Table body-->
            <tbody>
            <tr v-for="supplier in suppliers" :key="supplier">
              <td class="align-middle">
                <div class="d-flex align-items-center">
                  <div class="d-flex justify-content-start flex-column">
                    <span @click="showEditModal(supplier)" class="text-dark fw-bold text-hover-primary fs-6 cursor-pointer text-hover-primary">
                      {{supplier.name }}
                    </span>
                  </div>
                </div>
              </td>
              <td class="text-muted fw-semibold fs-6 align-middle">
                <span class="d-block fs-6">
                  {{ splitDate(supplier.CreatedAt) }}
                </span>
              </td>
              <td class="text-muted fw-semibold fs-6 align-middle">
                <span class="d-block fs-6">
                  {{ splitDate(supplier.UpdatedAt) }}
                </span>
              </td>
              <td class="align-middle">
                <span v-if="supplier.status === 1" class="badge badge-light-success">
                  Active
                </span>
              </td>
              <td class="border-0 align-middle">
                <div class="d-flex justify-content-end flex-shrink-0">
                  <button @click="showEditModal(supplier)" class="btn btn-icon btn-bg-light btn-active-color-primary btn-sm" data-v-5939e3ea=""><!--begin::Svg Icon | path: icons/duotune/art/art005.svg--><span class="svg-icon svg-icon-3" data-v-5939e3ea=""><svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" data-v-5939e3ea=""><path opacity="0.3" d="M21.4 8.35303L19.241 10.511L13.485 4.755L15.643 2.59595C16.0248 2.21423 16.5426 1.99988 17.0825 1.99988C17.6224 1.99988 18.1402 2.21423 18.522 2.59595L21.4 5.474C21.7817 5.85581 21.9962 6.37355 21.9962 6.91345C21.9962 7.45335 21.7817 7.97122 21.4 8.35303ZM3.68699 21.932L9.88699 19.865L4.13099 14.109L2.06399 20.309C1.98815 20.5354 1.97703 20.7787 2.03189 21.0111C2.08674 21.2436 2.2054 21.4561 2.37449 21.6248C2.54359 21.7934 2.75641 21.9115 2.989 21.9658C3.22158 22.0201 3.4647 22.0084 3.69099 21.932H3.68699Z" fill="currentColor" data-v-5939e3ea=""></path><path d="M5.574 21.3L3.692 21.928C3.46591 22.0032 3.22334 22.0141 2.99144 21.9594C2.75954 21.9046 2.54744 21.7864 2.3789 21.6179C2.21036 21.4495 2.09202 21.2375 2.03711 21.0056C1.9822 20.7737 1.99289 20.5312 2.06799 20.3051L2.696 18.422L5.574 21.3ZM4.13499 14.105L9.891 19.861L19.245 10.507L13.489 4.75098L4.13499 14.105Z" fill="currentColor" data-v-5939e3ea=""></path></svg></span><!--end::Svg Icon--></button>
                </div>
              </td>
            </tr>
            </tbody>
            <!--end::Table body-->
          </table>
          <EditSupplierModal ref="edit"/>
          <!--end::Table-->
        </div>
        <!--end::Table container-->
      </div>
      <!--begin::Body-->
    </div>
  </div>
</template>

<script>
import {mapGetters, mapActions} from 'vuex'
import EditSupplierModal from './modals/editSupplierModal.vue';
import timezones from "@/timezones.json";

export default {
  name: 'tableSuppliers',
  data() {
    return {
      timezones: timezones,
      check: false,
      req: true,
    }
  },

  methods: {
    ...mapActions({
      login: "auth/login",
      getAccess: "admin/getAccess"
    }),
    selectAllUsers() {
      if (this.check == true) {
        return this.checkedUsers = []
      }
      return [this.checkedUsers = this.filteredUsers, this.check == true]
    },
    showModal() {
      this.$refs.modal.show = true
    },
    showEditModal(supplier) {
      this.$refs.edit.openModal(supplier)
    },
    getSuppliers() {
      if (this.loggedIn && this.suppliers.length === 0 && this.req) {
        this.$store.dispatch("admin/getSuppliersMethod")
      }
      this.req = false
      return true
    },
    splitDate(date) {
      if (date === "0001-01-01T00:00:00Z" || !date) {
        return "-"
      }
      let userTimeZone = this.currentUser.timeZone
      let timeZone = this.timezones.find(item => item.text === userTimeZone);
      if(timeZone && timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc.split('-')[1])
        newDate.setTime(newDate.getTime() - (utc * 60 * 60 * 1000))
        return newDate.toLocaleString().split(",")[0] + " " + newDate.toLocaleString().split(",")[1]
      }
      if(timeZone && !timeZone.utc.includes('-')) {
        let newDate = new Date(date);
        let utc = Number(timeZone.utc)
        newDate.setTime(newDate.getTime() + (utc * 60 * 60 * 1000))
        return newDate.toLocaleString().split(",")[0] + " " + newDate.toLocaleString().split(",")[1]
      }
    },
  },
  computed: {
    ...mapGetters({
      suppliers: "admin/getSuppliers",
      loggedIn: "auth/isLoggedIn",
      getTokenHeader: "auth/getTokenHeader",
      currentUser: "auth/currentUser"
    })
  },
  components: {
    EditSupplierModal
  }
}
</script>

<style scoped>
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}
.table th:last-child{
  padding: 0.75rem 0.75rem !important;
}

.show-delete {
  transition: opacity 0.2s ease !important;
}
</style>