<template>
  <div class="d-flex flex-column flex-column-fluid container-fluid">
    <!--begin::Toolbar-->
    <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
      <!--begin::Page title-->
      <div class="page-title d-flex flex-column me-3">
        <!--begin::Title-->
        <h2 class="d-flex text-dark fw-bold my-1 fs-3">{{$t('cards.create')}}</h2>
        <!--end::Title-->
        <!--begin::Breadcrumb-->
        <ul class="breadcrumb breadcrumb-dot fw-semibold text-gray-600 fs-7 my-1">
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/" class="text-gray-600 text-hover-primary">{{ $t('common.crumb1') }}</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-600">
            <router-link to="/cards" class="text-gray-600 text-hover-primary">{{ $t('cards.title') }}</router-link>
          </li>
          <!--end::Item-->
          <!--begin::Item-->
          <li class="breadcrumb-item text-gray-500">{{ $t('cards.create') }}</li>
          <!--end::Item-->
        </ul>
        <!--end::Breadcrumb-->
      </div>
      <!--end::Page title-->
    </div>
    <!--end::Toolbar-->
    <!--begin::Post-->
    <div class="content flex-column-fluid" id="kt_content">
      <!--begin::Pricing card-->
      <div class="card" id="kt_pricing">
        <!--begin::Card body-->
        <div class="card-body p-5">
          <!--begin::Plans-->
          <div class="d-flex flex-column">
            <!--begin::Row-->
            <div class="row g-10">
              <!--begin::Col-->
              <div class="col-xl-3 col-sm-6">
                <div class="d-flex h-100 align-items-center">
                  <!--begin::Option-->
                  <div class="w-100 d-flex flex-column flex-center rounded-3 py-7 px-7" style="background-color: #d1e0ff6e">
                    <!--begin::Heading-->
                    <div class="mb-7 text-center">
                      <!--begin::Title-->
                      <h2 class="mb-5 fw-bolder" style="color: #1460ff">Facebook</h2>
                      <!--end::Title-->
                      <!--begin::Price-->
                      <div class="text-center">
                        <span class="mb-2" style="color: #1460ff">$</span>
                        <span class="fs-2x fw-bold" style="color: #1460ff">10</span>
                        <span class="fs-8 fw-semibold opacity-50">/
																	<span data-kt-element="period">Mon</span></span>
                      </div>
                      <!--end::Price-->
                    </div>
                    <!--end::Heading-->
                    <!--begin::Features-->
                    <div class="w-100 mb-10">
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1 pe-3">{{$t('rates.fbDescription')}} Facebook Ads</span>
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1 pe-3">{{$t('rates.cost')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen043.svg-->
                        <span class="text-info fw-semibold" style="color: rgb(20, 96, 255) !important;">
                          2.8%
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1 pe-3">{{$t('rates.commission')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen043.svg-->
                        <span class="text-info fw-semibold" style="color: rgb(20, 96, 255) !important;">
                          0.5%
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1">{{$t('rates.commissionFunds')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="svg-icon svg-icon-2">
																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																			<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="10" fill="currentColor"></rect>
																			<rect x="7" y="15.3137" width="12" height="2" rx="1" transform="rotate(-45 7 15.3137)" fill="currentColor"></rect>
																			<rect x="8.41422" y="7" width="12" height="2" rx="1" transform="rotate(45 8.41422 7)" fill="currentColor"></rect>
																		</svg>
																	</span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1">{{$t('rates.amountCards')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="text-info fw-semibold" style="color: rgb(20, 96, 255) !important;">
                         {{$t('rates.notLimited')}}
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1">{{$t('rates.issue')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="text-info fw-semibold" style="color: rgb(20, 96, 255) !important;">
                          2$
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center">
                        <span class="fw-semibold fs-7 text-gray-400 flex-grow-1">3D Secure</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="svg-icon svg-icon-2">
																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																			<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="10" fill="currentColor"></rect>
																			<rect x="7" y="15.3137" width="12" height="2" rx="1" transform="rotate(-45 7 15.3137)" fill="currentColor"></rect>
																			<rect x="8.41422" y="7" width="12" height="2" rx="1" transform="rotate(45 8.41422 7)" fill="currentColor"></rect>
																		</svg>
																	</span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                    </div>
                    <!--end::Features-->
                    <!--begin::Select-->
                    <button @click='showModal("facebook")' class="btn btn-sm btn-primary" style="background-color: #1460ff !important;">Select</button>
                    <!--end::Select-->
                    <CreateCardModal ref="facebook"/>
                  </div>
                  <!--end::Option-->
                </div>
              </div>
              <!--end::Col-->
              <!--begin::Col-->
              <div class="col-xl-3 col-sm-6">
                <div class="d-flex h-100 align-items-center">
                  <!--begin::Option-->
                  <div class="w-100 d-flex flex-column flex-center rounded-3 py-7 px-7" style="background-color: var(--kt-danger-light);">
                    <!--begin::Heading-->
                    <div class="mb-7 text-center">
                      <!--begin::Title-->
                      <h2 class="text-danger mb-5 fw-bolder">Advertising</h2>
                      <!--end::Title-->
                      <!--begin::Price-->
                      <div class="text-center">
                        <span class="mb-2 text-danger">$</span>
                        <span class="fs-2x fw-bold text-danger" >10</span>
                        <span class="fs-8 fw-semibold opacity-50">/
																	<span data-kt-element="period">Mon</span></span>
                      </div>
                      <!--end::Price-->
                    </div>
                    <!--end::Heading-->
                    <!--begin::Features-->
                    <div class="w-100 mb-10">
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1 pe-3">{{$t('rates.advDescription')}}</span>
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1 pe-3">{{$t('rates.cost')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen043.svg-->
                        <span class="text-danger fw-semibold">
                          5.8%
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1 pe-3">{{$t('rates.commission')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen043.svg-->
                        <span class="text-danger fw-semibold">
                          0%
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1">{{$t('rates.commissionFunds')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="svg-icon svg-icon-2">
																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																			<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="10" fill="currentColor"></rect>
																			<rect x="7" y="15.3137" width="12" height="2" rx="1" transform="rotate(-45 7 15.3137)" fill="currentColor"></rect>
																			<rect x="8.41422" y="7" width="12" height="2" rx="1" transform="rotate(45 8.41422 7)" fill="currentColor"></rect>
																		</svg>
																	</span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1">{{$t('rates.amountCards')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="text-danger fw-semibold">
                         {{$t('rates.notLimited')}}
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1">{{$t('rates.issue')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="text-danger fw-semibold">
                          2$
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center">
                        <span class="fw-semibold fs-7 text-gray-400 flex-grow-1">3D Secure</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="svg-icon svg-icon-2">
																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																			<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="10" fill="currentColor"></rect>
																			<rect x="7" y="15.3137" width="12" height="2" rx="1" transform="rotate(-45 7 15.3137)" fill="currentColor"></rect>
																			<rect x="8.41422" y="7" width="12" height="2" rx="1" transform="rotate(45 8.41422 7)" fill="currentColor"></rect>
																		</svg>
																	</span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                    </div>
                    <!--end::Features-->
                    <!--begin::Select-->
                    <button @click="showModal('advertising')" class="btn btn-sm btn-danger">Select</button>
                    <!--end::Select-->
                    <CreateCardModal ref="advertising"/>
                  </div>
                  <!--end::Option-->
                </div>
              </div>
              <!--end::Col-->
              <!--begin::Col-->
              <div class="col-xl-3 col-sm-6">
                <div class="d-flex h-100 align-items-center">
                  <!--begin::Option-->
                  <div class="w-100 d-flex flex-column flex-center rounded-3 py-7 px-7" style="background-color: var(--kt-info-light);">
                    <!--begin::Heading-->
                    <div class="mb-7 text-center">
                      <!--begin::Title-->
                      <h2 class="text-info mb-5 fw-bolder">Universal</h2>
                      <!--end::Title-->
                      <!--begin::Price-->
                      <div class="text-center">
                        <span class="mb-2 text-info">$</span>
                        <span class="fs-2x fw-bold text-info">15</span>
                        <span class="fs-8 fw-semibold opacity-50">/
																	<span data-kt-element="period">Mon</span></span>
                      </div>
                      <!--end::Price-->
                    </div>
                    <!--end::Heading-->
                    <!--begin::Features-->
                    <div class="w-100 mb-10">
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1 pe-3">{{$t('rates.uniDescription')}}</span>
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1 pe-3">{{$t('rates.cost')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen043.svg-->
                        <span class="text-info fw-semibold">
                          5.8%
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1 pe-3">{{$t('rates.commission')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen043.svg-->
                        <span class="text-info fw-semibold">
                          0%
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1">{{$t('rates.commissionFunds')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="svg-icon svg-icon-2">
																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																			<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="10" fill="currentColor"></rect>
																			<rect x="7" y="15.3137" width="12" height="2" rx="1" transform="rotate(-45 7 15.3137)" fill="currentColor"></rect>
																			<rect x="8.41422" y="7" width="12" height="2" rx="1" transform="rotate(45 8.41422 7)" fill="currentColor"></rect>
																		</svg>
																	</span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1">{{$t('rates.amountCards')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="text-info fw-semibold">
                         {{$t('rates.notLimited')}}
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center mb-5">
                        <span class="fw-semibold fs-7 text-gray-800 flex-grow-1">{{$t('rates.issue')}}</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="text-info fw-semibold">
                          1$
                        </span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                      <!--begin::Item-->
                      <div class="d-flex align-items-center">
                        <span class="fw-semibold fs-7 text-gray-400 flex-grow-1">3D Secure</span>
                        <!--begin::Svg Icon | path: icons/duotune/general/gen040.svg-->
                        <span class="svg-icon svg-icon-2">
																		<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																			<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="10" fill="currentColor"></rect>
																			<rect x="7" y="15.3137" width="12" height="2" rx="1" transform="rotate(-45 7 15.3137)" fill="currentColor"></rect>
																			<rect x="8.41422" y="7" width="12" height="2" rx="1" transform="rotate(45 8.41422 7)" fill="currentColor"></rect>
																		</svg>
																	</span>
                        <!--end::Svg Icon-->
                      </div>
                      <!--end::Item-->
                    </div>
                    <!--end::Features-->
                    <!--begin::Select-->
                    <button @click="showModal('universal')" class="btn btn-sm btn-info">Select</button>
                    <!--end::Select-->
                    <CreateCardModal ref="universal"/>
                  </div>
                  <!--end::Option-->
                </div>
              </div>
              <!--end::Col-->
              <!--begin::Col-->
              <div class="col-xl-3 col-sm-6">
                <div class="d-flex h-100 align-items-center">
                  <!--begin::Option-->
                  <div class="w-100 d-flex flex-column flex-center rounded-3 py-7 px-7 individual-background">
                    <!--begin::Heading-->
                    <div class="mb-7 text-center">
                      <!--begin::Title-->
                      <h2 class="mb-5 fw-bolder individual-color">Individual</h2>
                      <!--end::Title-->
                      <!--begin::Price-->
                      <div class="text-center">
                        <span class="mb-2 individual-color">$</span>
                        <span class="fs-2x fw-bold individual-color">?</span>
                        <span class="fs-8 fw-semibold opacity-50">/
																	<span data-kt-element="period">Mon</span></span>
                      </div>
                      <!--end::Price-->
                    </div>
                    <!--end::Heading-->

                    <!--begin::Select-->
                    <a href="#" class="btn btn-sm btn-dark">Select</a>
                    <!--end::Select-->
                  </div>
                  <!--end::Option-->
                </div>
              </div>
              <!--end::Col-->
            </div>
            <!--end::Row-->
          </div>
          <!--end::Plans-->
        </div>
        <!--end::Card body-->
      </div>
      <!--end::Pricing card-->
    </div>
    <!--end::Post-->
<!--    <CreateCardModal ref="universal" />-->
<!--    <CreateCardModal ref="advertising" />-->
  </div>
</template>

<script>
import {mapActions, mapGetters} from "vuex";
import CreateCardModal from "@/components/modals/createCardModal";
export default {
  name: "cardRates",
  data() {
    return {
      formFB: {
        userID: null,
        cardProvider: "Grats.Cards",
        cardType: "Facebook",
        balance: null,
      },
      formAdv: {
        userID: null,
        cardProvider: "Grats.Cards",
        cardType: "Facebook",
        balance: null,
      },
      formUni: {
        userID: null,
        cardProvider: "Grats.Cards",
        cardType: "Facebook",
        balance: null,
      }
    }
  },
  methods: {
    ...mapActions({
      createCardByTask: "users/createCardByTask"
    }),
    createCardFacebook() {
      this.formFB.userID = this.user.id
    },
    showModal(ref) {
      this.$refs[ref].openModal(ref)
    },
  },
  computed: {
    ...mapGetters({
      loggedIn: "auth/isLoggedIn",
      isLoaded: "auth/isLoaded",
      isTaskLoaded: "users/getTasksLoaded",
      user: "auth/currentUser"
    })
  },
  components: {
    CreateCardModal
  }
}
</script>

<style scoped>
.individual-color {
  background: linear-gradient(45deg, #160404 33%, #f40000 66%, #867fd4);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  color: #160404
}
/*.individual-background {*/
/*  background: rgb(22, 4, 4);;*/
/*  background: linear-gradient(45deg, rgba(22, 4, 4, 1) 33%, rgba(244, 0, 0, 1) 66%, #867fd4);*/
/*}*/
</style>