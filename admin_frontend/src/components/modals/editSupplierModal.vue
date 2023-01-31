<template>
  <!-- из шаблона -->
  <div v-if="show">
    <div :class="{show: this.isShow}" class="modal fade d-block" id="kt_modal_view_users" tabindex="-1"
         aria-hidden="true">
      <div class="modal-dialog modal-dialog-centered mw-550px" :style="{transform: this.transform}">
        <!--begin::Modal content-->
        <div class="modal-content">
          <!--begin::Modal header-->
          <div class="modal-header">
            <div class="text-center">
              <!--begin::Title-->
              <h1 class="text-dark fw-bolder">Change supplier</h1>
              <!--end::Title-->
            </div>
            <div @click="closeModal" class="btn btn-sm btn-icon btn-active-color-primary"
                 data-bs-dismiss="modal">
              <!--begin::Svg Icon | path: icons/duotune/arrows/arr061.svg-->
              <span class="svg-icon svg-icon-1">
                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                 xmlns="http://www.w3.org/2000/svg">
                                <rect opacity="0.5" x="6" y="17.3137" width="16" height="2" rx="1"
                                      transform="rotate(-45 6 17.3137)" fill="currentColor"/>
                                <rect x="7.41422" y="6" width="16" height="2" rx="1" transform="rotate(45 7.41422 6)"
                                      fill="currentColor"/>
                            </svg>
                        </span>
              <!--end::Svg Icon-->
            </div>
          </div>
          <div class="modal-body scroll-y">
            <div>
              <!--begin::List-->
              <div class="scroll-y me-n7 pe-7">
                <div class="card-body p-0">
                  <div class="card rounded-3">
                    <div class="card-body">
                      <form class="form w-100" id="registration"
                            @submit.prevent="checkForm">
                        <label class="required fs-6 fw-semibold mb-2" for="name">BINS</label>
                        <div v-for="bin in binsList" :key="bin" class="form-group fv-row mb-0">
                          <div class="d-flex align-items-center mb-5 position-relative">
                            <input type="number" class="form-control form-control-lg form-control-solid"
                                   v-model="bin.bin"
                                   id="name" name="name" placeholder="Enter BIN"
                                   required/>
                            <div class="form-check form-check-sm form-check-custom form-check-solid">
                              <input v-model="bin.active" class="form-check-input" type="checkbox"
                                     style="left: 7rem"/>
                            </div>
                            <Popper hover :content="isUniversalOutput(bin.isUniversal)">
                              <div class="form-check form-check-sm form-check-custom form-check-solid">
                                <input v-model="bin.isUniversal" class="form-check-input" type="checkbox" :checked="bin.isUniversal" style="margin-left: 0.5rem;"/>
                              </div>
                            </Popper>

                            <div class="position-relative ml-5">
                              <div @click="deleteBIN(bin)" class="btn btn-sm btn-icon btn-active-color-danger"
                                   data-bs-dismiss="modal">
                                <!--begin::Svg Icon | path: icons/duotune/arrows/arr061.svg-->
                                <span class="svg-icon svg-icon-1">
                            <svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                 xmlns="http://www.w3.org/2000/svg">
                                <rect opacity="0.5" x="6" y="17.3137" width="16" height="2" rx="1"
                                      transform="rotate(-45 6 17.3137)" fill="currentColor"/>
                                <rect x="7.41422" y="6" width="16" height="2" rx="1" transform="rotate(45 7.41422 6)"
                                      fill="currentColor"/>
                            </svg>
                        </span>
                                <!--end::Svg Icon-->
                              </div>
                            </div>
                          </div>
                        </div>
                        <button type="button" @click="addBIN" class="btn btn-primary" style="white-space: nowrap">Add BIN</button>
                      </form>
                    </div>
                  </div>
                </div>
              </div>
              <!--end::List-->
            </div>
            <!--end::Users-->
          </div>
          <!--end::Modal body-->
          <div class="modal-footer flex-center">
            <button @click="closeModal" class="btn btn-light me-3">Discard</button>
            <button @click="checkForm()" type="submit" id="kt_modal_new_target_submit"
                    class="btn btn-primary">
              <span class="indicator-label">Save</span>
              <span class="indicator-progress">Please wait...
									<span class="spinner-border spinner-border-sm align-middle ms-2"></span></span>
            </button>
          </div>
        </div>
        <!--end::Modal content-->
      </div>
    </div>

  </div>
</template>

<script>
import {defineComponent} from 'vue'
import {mapActions} from "vuex";
import 'vue-select/dist/vue-select.css';

const initialState = () => {
  return {
    show: false,
    isShow: false,
    transform: "translate(0, -50px)",
    supplierName: null,
    supplierStatus: null,
    supplierID: null,
    binsList: null,
    check: null,
    errors: {
      password: false,
      responseError: null
    },
    form: {
      id: 0,
      status: 1,
      bins: null
    }
  }
}

export default defineComponent({
  name: "editSupplierModal",
  data() {
    return initialState();
  },
  methods: {
    ...mapActions({
      updateSupplier: "admin/updateSupplier"
    }),
    isUniversalOutput(universal) {
      if(universal) {
        return "The card with this bin is universal"
      }
      return "The card with this bin is advertising"
    },
    addBIN() {
      this.binsList.push({"bin": null, "active": true})
    },
    deleteBIN(bin) {
      if (this.binsList.indexOf(bin) !== -1) {
        this.binsList.splice(this.binsList.indexOf(bin), 1);
      }
    },
    openModal(supplier) {
      this.show = true
      setTimeout(()=> {
        this.isShow = this.show
        this.transform = "none"
      },200)
      this.supplierID = supplier.ID
      this.supplierStatus = supplier.status
      this.binsList = JSON.parse(supplier.bins)
    },
    closeModal: function () {
      this.isShow = false
      this.transform = "translate(0, -50px)"
      setTimeout(()=> {
        this.show = false
      },200)
    },
    checkForm: function (e) {
      this.form.id = this.supplierID
      this.form.status = this.supplierStatus
      this.form.bins = JSON.stringify(this.binsList)
          this.updateSupplier(this.form)
              .then(() => {
                Object.assign(this.$data, initialState());
                this.closeModal()
                this.$swal({
                  title: "You have successfully added a bin",
                  confirmButtonColor: "#50cd89",
                  type: "alert message",
                  icon: "success",
                })
              })
              .catch(error => {
                console.log(error)
                this.errors.responseError = error
                let errorTitle = error
                this.$swal({
                  title: errorTitle,
                  type: "alert message",
                  confirmButtonColor: "#f1416c",
                  icon: "error",
                })
              });
      if (this.errors.responseError) {
        e.preventDefault();
      }
    },
  }
})
</script>

<style>
.table.gs-0 th:first-child,
.table.gs-0 td:first-child {
  padding-left: 1rem !important;
}

.modal {
  background: rgba(0, 0, 0, .5);
}
</style>