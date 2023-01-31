<template>
  <div class="d-flex flex-column flex-column-fluid container-fluid">
    <!--begin::Toolbar-->
    <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
      <!--begin::Page title-->
      <div class="page-title d-flex flex-column me-3">
        <!--begin::Title-->
        <h1 v-if="getDataInfo()" class="d-flex text-dark fw-bold my-1 fs-3">My Balance</h1>
        <!--end::Title-->
      </div>
      <!--end::Page title-->
      <!--begin::Actions-->
      <div class="d-flex align-items-center py-2 py-md-1 d-none">
        <!--begin::Wrapper-->
        <div class="me-3">
          <!--begin::Menu-->
          <a href="#" class="btn btn-light fw-bold" data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end">
            <!--begin::Svg Icon | path: icons/duotune/general/gen031.svg-->
            <span class="svg-icon svg-icon-5 svg-icon-gray-500 me-1">
											<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
												<path d="M19.0759 3H4.72777C3.95892 3 3.47768 3.83148 3.86067 4.49814L8.56967 12.6949C9.17923 13.7559 9.5 14.9582 9.5 16.1819V19.5072C9.5 20.2189 10.2223 20.7028 10.8805 20.432L13.8805 19.1977C14.2553 19.0435 14.5 18.6783 14.5 18.273V13.8372C14.5 12.8089 14.8171 11.8056 15.408 10.964L19.8943 4.57465C20.3596 3.912 19.8856 3 19.0759 3Z" fill="currentColor"></path>
											</svg>
										</span>
            <!--end::Svg Icon-->Filter</a>
          <!--begin::Menu 1-->
          <div class="menu menu-sub menu-sub-dropdown w-250px w-md-300px" data-kt-menu="true" id="kt_menu_62cfb082372eb">
            <!--begin::Header-->
            <div class="px-7 py-5">
              <div class="fs-5 text-dark fw-bold">Filter Options</div>
            </div>
            <!--end::Header-->
            <!--begin::Menu separator-->
            <div class="separator border-gray-200"></div>
            <!--end::Menu separator-->
            <!--begin::Form-->
            <div class="px-7 py-5">
              <!--begin::Input group-->
              <div class="mb-10">
                <!--begin::Label-->
                <label class="form-label fw-semibold">Status:</label>
                <!--end::Label-->
                <!--begin::Input-->
                <div>
                  <select class="form-select form-select-solid select2-hidden-accessible" data-kt-select2="true" data-placeholder="Select option" data-dropdown-parent="#kt_menu_62cfb082372eb" data-allow-clear="true" data-select2-id="select2-data-7-uc6q" tabindex="-1" aria-hidden="true" data-kt-initialized="1">
                    <option data-select2-id="select2-data-9-2p9f"></option>
                    <option value="1">Approved</option>
                    <option value="2">Pending</option>
                    <option value="2">In Process</option>
                    <option value="2">Rejected</option>
                  </select><span class="select2 select2-container select2-container--bootstrap5" dir="ltr" data-select2-id="select2-data-8-z1lm" style="width: 100%;"><span class="selection"><span class="select2-selection select2-selection--single form-select form-select-solid" role="combobox" aria-haspopup="true" aria-expanded="false" tabindex="0" aria-disabled="false" aria-labelledby="select2-jmzy-container" aria-controls="select2-jmzy-container"><span class="select2-selection__rendered" id="select2-jmzy-container" role="textbox" aria-readonly="true" title="Select option"><span class="select2-selection__placeholder">Select option</span></span><span class="select2-selection__arrow" role="presentation"><b role="presentation"></b></span></span></span><span class="dropdown-wrapper" aria-hidden="true"></span></span>
                </div>
                <!--end::Input-->
              </div>
              <!--end::Input group-->
              <!--begin::Input group-->
              <div class="mb-10">
                <!--begin::Label-->
                <label class="form-label fw-semibold">Member Type:</label>
                <!--end::Label-->
                <!--begin::Options-->
                <div class="d-flex">
                  <!--begin::Options-->
                  <label class="form-check form-check-sm form-check-custom form-check-solid me-5">
                    <input class="form-check-input" type="checkbox" value="1">
                    <span class="form-check-label">Author</span>
                  </label>
                  <!--end::Options-->
                  <!--begin::Options-->
                  <label class="form-check form-check-sm form-check-custom form-check-solid">
                    <input class="form-check-input" type="checkbox" value="2" checked="checked">
                    <span class="form-check-label">Customer</span>
                  </label>
                  <!--end::Options-->
                </div>
                <!--end::Options-->
              </div>
              <!--end::Input group-->
              <!--begin::Input group-->
              <div class="mb-10">
                <!--begin::Label-->
                <label class="form-label fw-semibold">Notifications:</label>
                <!--end::Label-->
                <!--begin::Switch-->
                <div class="form-check form-switch form-switch-sm form-check-custom form-check-solid">
                  <input class="form-check-input" type="checkbox" value="" name="notifications" checked="checked">
                  <label class="form-check-label">Enabled</label>
                </div>
                <!--end::Switch-->
              </div>
              <!--end::Input group-->
              <!--begin::Actions-->
              <div class="d-flex justify-content-end">
                <button type="reset" class="btn btn-sm btn-light btn-active-light-primary me-2" data-kt-menu-dismiss="true">Reset</button>
                <button type="submit" class="btn btn-sm btn-primary" data-kt-menu-dismiss="true">Apply</button>
              </div>
              <!--end::Actions-->
            </div>
            <!--end::Form-->
          </div>
          <!--end::Menu 1-->
          <!--end::Menu-->
        </div>
        <!--end::Wrapper-->
        <!--begin::Button-->
        <a href="#" class="btn btn-dark fw-bold" data-bs-toggle="modal" data-bs-target="#kt_modal_create_app" id="kt_toolbar_primary_button">Create</a>
        <!--end::Button-->
      </div>
      <!--end::Actions-->
    </div>
    <!--end::Toolbar-->
    <!--begin::Post-->
    <div class="content flex-column-fluid" id="kt_content">
      <!--begin::Row-->
      <div class="row g-5 g-xl-10 mb-5 mb-xl-10">
        <!--begin::Col-->
        <div class="col-xxl-12">
          <!--begin::Row-->
          <div class="row g-5 g-xl-12">
            <!--begin::Col-->
            <div class="col-md-3">
              <!--begin::Card widget 11-->
              <div class="card card-flush h-xl-100" style="background-color: #F6E5CA">
                <!--begin::Header-->
                <div class="card-header flex-nowrap pt-5">
                  <!--begin::Title-->
                  <h3 class="card-title align-items-start flex-column">
                    <span class="card-label fw-bold fs-4 text-gray-800">Dollar</span>
                    <span class="mt-1 fw-semibold fs-7" style="display: none">36,668 USD for 1 BTC</span>
                  </h3>
                  <!--end::Title-->
                  <!--begin::Toolbar-->
                  <div class="card-toolbar">
                    <!--begin::Menu-->
                    <button class="btn btn-icon justify-content-end" data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end" data-kt-menu-overflow="true" style="color:">
                      <!--begin::Svg Icon | path: icons/duotune/general/gen023.svg-->
                      <span class="svg-icon svg-icon-1">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="4" fill="currentColor"></rect>
																		<rect x="11" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																		<rect x="15" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																		<rect x="7" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																	</svg>
																</span>
                      <!--end::Svg Icon-->
                    </button>
                    <!--begin::Menu 2-->
                    <div class="menu menu-sub menu-sub-dropdown menu-column menu-rounded menu-gray-800 menu-state-bg-light-primary fw-semibold w-200px" data-kt-menu="true">
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <div class="menu-content fs-6 text-dark fw-bold px-3 py-4">Quick Actions</div>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu separator-->
                      <div class="separator mb-3 opacity-75"></div>
                      <!--end::Menu separator-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Ticket</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Customer</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3" data-kt-menu-trigger="hover" data-kt-menu-placement="right-start">
                        <!--begin::Menu item-->
                        <a href="#" class="menu-link px-3">
                          <span class="menu-title">New Group</span>
                          <span class="menu-arrow"></span>
                        </a>
                        <!--end::Menu item-->
                        <!--begin::Menu sub-->
                        <div class="menu-sub menu-sub-dropdown w-175px py-4">
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Admin Group</a>
                          </div>
                          <!--end::Menu item-->
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Staff Group</a>
                          </div>
                          <!--end::Menu item-->
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Member Group</a>
                          </div>
                          <!--end::Menu item-->
                        </div>
                        <!--end::Menu sub-->
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Contact</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu separator-->
                      <div class="separator mt-3 opacity-75"></div>
                      <!--end::Menu separator-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <div class="menu-content px-3 py-3">
                          <a class="btn btn-primary btn-sm px-4" href="#">Generate Reports</a>
                        </div>
                      </div>
                      <!--end::Menu item-->
                    </div>
                    <!--end::Menu 2-->
                    <!--end::Menu-->
                  </div>
                  <!--end::Toolbar-->
                </div>
                <!--end::Header-->
                <!--begin::Body-->
                <div class="card-body text-center pt-5">
                  <!--begin::Image-->
                  <div class="h-125px mb-5 d-none">
                    <svg xmlns="http://www.w3.org/2000/svg" width="130" height="130" viewBox="0 0 130 130" fill="none">
                      <path d="M65 4.0625C52.9477 4.06246 41.1661 7.63636 31.1449 14.3322C21.1238 21.0281 13.3133 30.5453 8.70109 41.6802C4.08888 52.815 2.88213 65.0675 5.23345 76.8882C7.58477 88.709 13.3885 99.567 21.9108 108.089C27.569 113.749 34.2868 118.24 41.6805 121.304C49.0742 124.367 56.999 125.945 65.0024 125.945C73.0058 125.946 80.9309 124.37 88.3252 121.308C95.7195 118.246 102.438 113.757 108.097 108.097C113.757 102.438 118.246 95.7195 121.308 88.3252C124.37 80.9309 125.946 73.0058 125.945 65.0024C125.945 56.999 124.367 49.0742 121.304 41.6805C118.24 34.2868 113.749 27.569 108.089 21.9108C102.444 16.2345 95.7287 11.7342 88.3324 8.67054C80.9362 5.60687 73.0057 4.04064 65 4.0625V4.0625ZM105.217 105.217C99.9355 110.499 93.6656 114.689 86.765 117.548C79.8644 120.407 72.4682 121.879 64.9988 121.879C57.5295 121.879 50.1331 120.408 43.2322 117.55C36.3314 114.692 30.0611 110.502 24.7794 105.221C19.4977 99.939 15.3082 93.6687 12.4499 86.7678C9.59169 79.8669 8.12076 72.4706 8.12113 65.0012C8.1215 57.5318 9.59317 50.1356 12.4521 43.235C15.311 36.3344 19.5012 30.0645 24.7834 24.7834C30.0645 19.5012 36.3344 15.311 43.235 12.4521C50.1356 9.59315 57.5318 8.12148 65.0012 8.1211C72.4706 8.12073 79.8669 9.59166 86.7678 12.4499C93.6687 15.3081 99.939 19.4977 105.221 24.7794C110.502 30.061 114.692 36.3313 117.55 43.2322C120.408 50.1331 121.879 57.5294 121.879 64.9988C121.879 72.4682 120.407 79.8644 117.548 86.765C114.689 93.6656 110.499 99.9355 105.217 105.217V105.217Z" fill="#E0C6A0"/>
                      <path d="M54.8438 60.9375H70.0499C72.4562 60.9398 74.8091 60.2275 76.8101 58.8909L76.8104 58.8906C78.5908 57.7037 79.9421 55.9754 80.6646 53.9612C81.387 51.947 81.4423 49.7539 80.8222 47.7059C80.2021 45.6579 78.9396 43.8637 77.2212 42.5886C75.5028 41.3135 73.4198 40.625 71.28 40.625H54.8438C54.305 40.625 53.7884 40.839 53.4075 41.22C53.0265 41.6009 52.8125 42.1175 52.8125 42.6563V58.9063C52.8125 59.445 53.0265 59.9616 53.4075 60.3425C53.7884 60.7235 54.305 60.9375 54.8438 60.9375ZM56.875 44.6875H71.28C72.5479 44.6875 73.7821 45.0955 74.8003 45.851C75.8185 46.6066 76.5666 47.6697 76.934 48.8832C77.3014 50.0967 77.2687 51.3962 76.8406 52.5896C76.4126 53.7831 75.6119 54.8072 74.5569 55.5105C73.2228 56.4016 71.6542 56.8765 70.0499 56.875H56.875V44.6875Z" fill="#E0C6A0"/>
                      <path d="M54.8438 89.375H74.4935C76.9307 89.375 79.2965 88.5521 81.2076 87.0397C83.1187 85.5272 84.4631 83.4137 85.0231 81.0417C85.583 78.6697 85.3257 76.1782 84.2927 73.9707C83.2598 71.7632 81.5117 69.9692 79.3318 68.8793C76.9 67.6671 74.2205 67.0346 71.5034 67.0312H54.8438C54.305 67.0313 53.7884 67.2453 53.4075 67.6262C53.0265 68.0071 52.8125 68.5238 52.8125 69.0625V87.3438C52.8125 87.8825 53.0265 88.3991 53.4075 88.78C53.7884 89.161 54.305 89.375 54.8438 89.375ZM56.875 71.0938H71.5034C73.5899 71.0963 75.6476 71.582 77.515 72.5128C78.8764 73.1935 79.9681 74.3138 80.6131 75.6924C81.2582 77.071 81.4189 78.627 81.0693 80.1083C80.7196 81.5896 79.88 82.9095 78.6865 83.8541C77.493 84.7986 76.0155 85.3125 74.4935 85.3125H56.875V71.0938Z" fill="#E0C6A0"/>
                      <path d="M28.0135 36.4541L24.799 33.9701C23.4136 35.7633 22.149 37.6467 21.0137 39.6079L24.5302 41.6421C25.5749 39.8374 26.7386 38.1042 28.0135 36.4541V36.4541Z" fill="#E0C6A0"/>
                      <path d="M53.3188 19.7529L52.3062 15.8188C50.1459 16.3753 48.0248 17.0741 45.9568 17.9105L47.4807 21.6765C49.3823 20.9073 51.3325 20.2647 53.3188 19.7529V19.7529Z" fill="#E0C6A0"/>
                      <path d="M65 14.2188C62.8789 14.2188 60.7598 14.35 58.6549 14.6115L59.1577 18.6428C61.0959 18.402 63.047 18.2813 65 18.2812C90.7609 18.2813 111.719 39.2392 111.719 65C111.719 90.7608 90.7609 111.719 65 111.719C39.2392 111.719 18.2813 90.7608 18.2813 65C18.2674 58.5641 19.593 52.1958 22.1737 46.3L18.4514 44.6724C15.6457 51.0813 14.2043 58.0039 14.2188 65C14.2188 93.0009 36.9992 115.781 65 115.781C93.0009 115.781 115.781 93.0009 115.781 65C115.781 36.9991 93.0009 14.2188 65 14.2188Z" fill="#E0C6A0"/>
                      <path d="M40.625 97.5H48.75V103.594C48.75 104.132 48.964 104.649 49.345 105.03C49.7259 105.411 50.2425 105.625 50.7812 105.625H58.9062C59.445 105.625 59.9616 105.411 60.3425 105.03C60.7235 104.649 60.9375 104.132 60.9375 103.594V97.5H65V103.594C65 104.132 65.214 104.649 65.595 105.03C65.9759 105.411 66.4925 105.625 67.0312 105.625H75.1562C75.695 105.625 76.2116 105.411 76.5925 105.03C76.9735 104.649 77.1875 104.132 77.1875 103.594V97.4903C81.2394 97.3594 85.1417 95.9284 88.3178 93.409C91.4938 90.8895 93.7752 87.4152 94.8247 83.4994C95.8742 79.5836 95.6362 75.434 94.1459 71.6639C92.6556 67.8937 89.9921 64.703 86.5488 62.5631C88.7255 60.2542 90.2468 57.4063 90.9556 54.3132C91.6644 51.2202 91.5352 47.9941 90.5812 44.9677C89.6272 41.9412 87.8831 39.2242 85.5286 37.0967C83.1742 34.9693 80.2948 33.5086 77.1875 32.8652V26.4063C77.1875 25.8675 76.9735 25.3509 76.5925 24.97C76.2116 24.589 75.695 24.375 75.1562 24.375H67.0312C66.4925 24.375 65.9759 24.589 65.595 24.97C65.214 25.3509 65 25.8675 65 26.4063V32.5H60.9375V26.4063C60.9375 25.8675 60.7235 25.3509 60.3425 24.97C59.9616 24.589 59.445 24.375 58.9062 24.375H50.7812C50.2425 24.375 49.7259 24.589 49.345 24.97C48.964 25.3509 48.75 25.8675 48.75 26.4063V32.5H40.625C40.0863 32.5 39.5696 32.714 39.1887 33.095C38.8078 33.4759 38.5938 33.9925 38.5938 34.5313V42.6563C38.5938 43.195 38.8078 43.7116 39.1887 44.0925C39.5696 44.4735 40.0863 44.6875 40.625 44.6875H42.6562V85.3125H40.625C40.0863 85.3125 39.5696 85.5265 39.1887 85.9075C38.8078 86.2884 38.5938 86.805 38.5938 87.3438V95.4688C38.5938 96.0075 38.8078 96.5241 39.1887 96.905C39.5696 97.286 40.0863 97.5 40.625 97.5V97.5ZM42.6562 89.375H44.6875C45.2262 89.375 45.7429 89.161 46.1238 88.78C46.5047 88.3991 46.7187 87.8825 46.7188 87.3438V42.6563C46.7187 42.1175 46.5047 41.6009 46.1238 41.22C45.7429 40.839 45.2262 40.625 44.6875 40.625H42.6562V36.5625H50.7812C51.32 36.5625 51.8366 36.3485 52.2175 35.9675C52.5985 35.5866 52.8125 35.07 52.8125 34.5313V28.4375H56.875V34.5313C56.875 35.07 57.089 35.5866 57.47 35.9675C57.8509 36.3485 58.3675 36.5625 58.9062 36.5625H67.0312C67.57 36.5625 68.0866 36.3485 68.4675 35.9675C68.8485 35.5866 69.0625 35.07 69.0625 34.5313V28.4375H73.125V34.6101C73.125 35.1141 73.3123 35.6 73.6506 35.9736C73.9888 36.3472 74.4539 36.5816 74.9554 36.6314C77.5478 36.8926 80.0129 37.884 82.0642 39.4904C84.1156 41.0968 85.669 43.2523 86.5441 45.7065C87.4191 48.1607 87.5798 50.8128 87.0075 53.3546C86.4352 55.8965 85.1534 58.2238 83.311 60.0662L81.8449 61.5323C81.6266 61.7506 81.4609 62.0158 81.3605 62.3077C81.26 62.5996 81.2274 62.9106 81.2651 63.2169C81.3028 63.5233 81.4099 63.8171 81.5781 64.0759C81.7464 64.3347 81.9714 64.5518 82.2361 64.7106L84.2066 65.8929C86.9875 67.5614 89.1447 70.0955 90.3479 73.1071C91.5511 76.1186 91.7339 79.4416 90.8682 82.5669C90.0025 85.6923 88.1362 88.4477 85.555 90.411C82.9739 92.3744 79.8203 93.4375 76.5773 93.4375H75.1562C74.6175 93.4375 74.1009 93.6515 73.72 94.0325C73.339 94.4134 73.125 94.93 73.125 95.4688V101.563H69.0625V95.4688C69.0625 94.93 68.8485 94.4134 68.4675 94.0325C68.0866 93.6515 67.57 93.4375 67.0312 93.4375H58.9062C58.3675 93.4375 57.8509 93.6515 57.47 94.0325C57.089 94.4134 56.875 94.93 56.875 95.4688V101.563H52.8125V95.4688C52.8125 94.93 52.5985 94.4134 52.2175 94.0325C51.8366 93.6515 51.32 93.4375 50.7812 93.4375H42.6562V89.375Z" fill="#E0C6A0"/>
                    </svg>
                  </div>
                  <!--end::Image-->
                  <!--begin::Section-->
                  <div class="text-start">
                    <span v-if="this.userinfo.USDBalance !== 0" class="d-block fw-bold fs-1 text-gray-800">{{this.userinfo.USDBalance}} $</span>
                    <span v-else class="d-block fw-bold fs-1 text-gray-800">0 $</span>
                    <span class="mt-1 fw-semibold fs-3 d-none" style="color:">19,335,45 USD</span>
                  </div>
                  <!--end::Section-->
                </div>
                <!--end::Body-->
              </div>
              <!--end::Card widget 11-->
            </div>
            <!--end::Col-->
            <!--begin::Col-->
            <div class="col-md-3">
              <!--begin::Card widget 11-->
              <div class="card card-flush h-xl-100" style="background-color: #F3D6EF">
                <!--begin::Header-->
                <div class="card-header flex-nowrap pt-5">
                  <!--begin::Title-->
                  <h3 class="card-title align-items-start flex-column">
                    <span class="card-label fw-bold fs-4 text-gray-800">Euro</span>
                    <span class="mt-1 fw-semibold fs-7 d-none" style="color:">325,035 USD for 1 ETH</span>
                  </h3>
                  <!--end::Title-->
                  <!--begin::Toolbar-->
                  <div class="card-toolbar">
                    <!--begin::Menu-->
                    <button class="btn btn-icon justify-content-end" data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end" data-kt-menu-overflow="true" style="color:">
                      <!--begin::Svg Icon | path: icons/duotune/general/gen023.svg-->
                      <span class="svg-icon svg-icon-1">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="4" fill="currentColor"></rect>
																		<rect x="11" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																		<rect x="15" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																		<rect x="7" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																	</svg>
																</span>
                      <!--end::Svg Icon-->
                    </button>
                    <!--begin::Menu 2-->
                    <div class="menu menu-sub menu-sub-dropdown menu-column menu-rounded menu-gray-800 menu-state-bg-light-primary fw-semibold w-200px" data-kt-menu="true">
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <div class="menu-content fs-6 text-dark fw-bold px-3 py-4">Quick Actions</div>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu separator-->
                      <div class="separator mb-3 opacity-75"></div>
                      <!--end::Menu separator-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Ticket</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Customer</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3" data-kt-menu-trigger="hover" data-kt-menu-placement="right-start">
                        <!--begin::Menu item-->
                        <a href="#" class="menu-link px-3">
                          <span class="menu-title">New Group</span>
                          <span class="menu-arrow"></span>
                        </a>
                        <!--end::Menu item-->
                        <!--begin::Menu sub-->
                        <div class="menu-sub menu-sub-dropdown w-175px py-4">
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Admin Group</a>
                          </div>
                          <!--end::Menu item-->
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Staff Group</a>
                          </div>
                          <!--end::Menu item-->
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Member Group</a>
                          </div>
                          <!--end::Menu item-->
                        </div>
                        <!--end::Menu sub-->
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Contact</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu separator-->
                      <div class="separator mt-3 opacity-75"></div>
                      <!--end::Menu separator-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <div class="menu-content px-3 py-3">
                          <a class="btn btn-primary btn-sm px-4" href="#">Generate Reports</a>
                        </div>
                      </div>
                      <!--end::Menu item-->
                    </div>
                    <!--end::Menu 2-->
                    <!--end::Menu-->
                  </div>
                  <!--end::Toolbar-->
                </div>
                <!--end::Header-->
                <!--begin::Body-->
                <div class="card-body text-center pt-5">
                  <!--begin::Image-->
<!--                  <img src="assets/media/svg/shapes/ethereum.svg" class="h-125px mb-5 d-none" alt="">-->
                  <!--end::Image-->
                  <!--begin::Section-->
                  <div class="text-start">
                    <span v-if="this.userinfo.EURBalance !== 0" class="d-block fw-bold fs-1 text-gray-800">{{this.userinfo.EURBalance}} €</span>
                    <span v-else class="d-block fw-bold fs-1 text-gray-800">0 €</span>
                    <span class="mt-1 fw-semibold fs-3 d-none">7,336,00 USD</span>
                  </div>
                  <!--end::Section-->
                </div>
                <!--end::Body-->
              </div>
              <!--end::Card widget 11-->
            </div>
            <!--end::Col-->
            <!--begin::Col-->
            <div class="col-md-3">
              <!--begin::Card widget 11-->
              <div class="card card-flush h-xl-100" style="background-color: #BFDDE3">
                <!--begin::Header-->
                <div class="card-header flex-nowrap pt-5">
                  <!--begin::Title-->
                  <h3 class="card-title align-items-start flex-column">
                    <span class="card-label fw-bold fs-4 text-gray-800">USDT</span>
                    <span class="mt-1 fw-semibold fs-7 d-none" style="color:">0.12,045 USD for 1 DOGE</span>
                  </h3>
                  <!--end::Title-->
                  <!--begin::Toolbar-->
                  <div class="card-toolbar">
                    <!--begin::Menu-->
                    <button class="btn btn-icon justify-content-end" data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end" data-kt-menu-overflow="true" style="color:">
                      <!--begin::Svg Icon | path: icons/duotune/general/gen023.svg-->
                      <span class="svg-icon svg-icon-1">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="4" fill="currentColor"></rect>
																		<rect x="11" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																		<rect x="15" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																		<rect x="7" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																	</svg>
																</span>
                      <!--end::Svg Icon-->
                    </button>
                    <!--begin::Menu 2-->
                    <div class="menu menu-sub menu-sub-dropdown menu-column menu-rounded menu-gray-800 menu-state-bg-light-primary fw-semibold w-200px" data-kt-menu="true">
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <div class="menu-content fs-6 text-dark fw-bold px-3 py-4">Quick Actions</div>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu separator-->
                      <div class="separator mb-3 opacity-75"></div>
                      <!--end::Menu separator-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Ticket</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Customer</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3" data-kt-menu-trigger="hover" data-kt-menu-placement="right-start">
                        <!--begin::Menu item-->
                        <a href="#" class="menu-link px-3">
                          <span class="menu-title">New Group</span>
                          <span class="menu-arrow"></span>
                        </a>
                        <!--end::Menu item-->
                        <!--begin::Menu sub-->
                        <div class="menu-sub menu-sub-dropdown w-175px py-4">
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Admin Group</a>
                          </div>
                          <!--end::Menu item-->
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Staff Group</a>
                          </div>
                          <!--end::Menu item-->
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Member Group</a>
                          </div>
                          <!--end::Menu item-->
                        </div>
                        <!--end::Menu sub-->
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Contact</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu separator-->
                      <div class="separator mt-3 opacity-75"></div>
                      <!--end::Menu separator-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <div class="menu-content px-3 py-3">
                          <a class="btn btn-primary btn-sm px-4" href="#">Generate Reports</a>
                        </div>
                      </div>
                      <!--end::Menu item-->
                    </div>
                    <!--end::Menu 2-->
                    <!--end::Menu-->
                  </div>
                  <!--end::Toolbar-->
                </div>
                <!--end::Header-->
                <!--begin::Body-->
                <div class="card-body text-center pt-5">
                  <!--begin::Image-->
<!--                  <img src="assets/media/svg/shapes/dogecoin.svg" class="h-125px mb-5 d-none" alt="">-->
                  <!--end::Image-->
                  <!--begin::Section-->
                  <div class="text-start">
                    <span v-if="this.userinfo.USDTBalance !== 0" class="d-block fw-bold fs-1 text-gray-800">{{this.userinfo.USDTBalance}} ₮</span>
                    <span v-else class="d-block fw-bold fs-1 text-gray-800">0 ₮</span>
                    <span class="mt-1 fw-semibold fs-3 d-none" style="color:">503,005,56 USD</span>
                  </div>
                  <!--end::Section-->
                </div>
                <!--end::Body-->
              </div>
              <!--end::Card widget 11-->
            </div>
            <!--end::Col-->
            <!--begin::Col-->
            <div class="col-md-3">
              <!--begin::Card widget 11-->
              <div class="card card-flush h-xl-100" style="background-color: #BFDDE3">
                <!--begin::Header-->
                <div class="card-header flex-nowrap pt-5">
                  <!--begin::Title-->
                  <h3 class="card-title align-items-start flex-column">
                    <span class="card-label fw-bold fs-4 text-gray-800">Bitcoin</span>
                    <span class="mt-1 fw-semibold fs-7 d-none" style="color:">0.12,045 USD for 1 DOGE</span>
                  </h3>
                  <!--end::Title-->
                  <!--begin::Toolbar-->
                  <div class="card-toolbar">
                    <!--begin::Menu-->
                    <button class="btn btn-icon justify-content-end" data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end" data-kt-menu-overflow="true" style="color:">
                      <!--begin::Svg Icon | path: icons/duotune/general/gen023.svg-->
                      <span class="svg-icon svg-icon-1">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<rect opacity="0.3" x="2" y="2" width="20" height="20" rx="4" fill="currentColor"></rect>
																		<rect x="11" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																		<rect x="15" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																		<rect x="7" y="11" width="2.6" height="2.6" rx="1.3" fill="currentColor"></rect>
																	</svg>
																</span>
                      <!--end::Svg Icon-->
                    </button>
                    <!--begin::Menu 2-->
                    <div class="menu menu-sub menu-sub-dropdown menu-column menu-rounded menu-gray-800 menu-state-bg-light-primary fw-semibold w-200px" data-kt-menu="true">
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <div class="menu-content fs-6 text-dark fw-bold px-3 py-4">Quick Actions</div>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu separator-->
                      <div class="separator mb-3 opacity-75"></div>
                      <!--end::Menu separator-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Ticket</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Customer</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3" data-kt-menu-trigger="hover" data-kt-menu-placement="right-start">
                        <!--begin::Menu item-->
                        <a href="#" class="menu-link px-3">
                          <span class="menu-title">New Group</span>
                          <span class="menu-arrow"></span>
                        </a>
                        <!--end::Menu item-->
                        <!--begin::Menu sub-->
                        <div class="menu-sub menu-sub-dropdown w-175px py-4">
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Admin Group</a>
                          </div>
                          <!--end::Menu item-->
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Staff Group</a>
                          </div>
                          <!--end::Menu item-->
                          <!--begin::Menu item-->
                          <div class="menu-item px-3">
                            <a href="#" class="menu-link px-3">Member Group</a>
                          </div>
                          <!--end::Menu item-->
                        </div>
                        <!--end::Menu sub-->
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <a href="#" class="menu-link px-3">New Contact</a>
                      </div>
                      <!--end::Menu item-->
                      <!--begin::Menu separator-->
                      <div class="separator mt-3 opacity-75"></div>
                      <!--end::Menu separator-->
                      <!--begin::Menu item-->
                      <div class="menu-item px-3">
                        <div class="menu-content px-3 py-3">
                          <a class="btn btn-primary btn-sm px-4" href="#">Generate Reports</a>
                        </div>
                      </div>
                      <!--end::Menu item-->
                    </div>
                    <!--end::Menu 2-->
                    <!--end::Menu-->
                  </div>
                  <!--end::Toolbar-->
                </div>
                <!--end::Header-->
                <!--begin::Body-->
                <div class="card-body text-center pt-5">
                  <!--begin::Image-->
<!--                  <img src="assets/media/svg/shapes/dogecoin.svg" class="h-125px mb-5 d-none" alt="">-->
                  <!--end::Image-->
                  <!--begin::Section-->
                  <div class="text-start">
                    <span v-if="this.userinfo.BTCBalance !== 0" class="d-block fw-bold fs-1 text-gray-800">{{this.userinfo.BTCBalance}} ₿</span>
                    <span v-else class="d-block fw-bold fs-1 text-gray-800">0 ₿</span>
                    <span class="mt-1 fw-semibold fs-3 d-none" style="color:">503,005,56 USD</span>
                  </div>
                  <!--end::Section-->
                </div>
                <!--end::Body-->
              </div>
              <!--end::Card widget 11-->
            </div>
            <!--end::Col-->
          </div>
          <!--end::Row-->
        </div>
        <!--end::Col-->
        <!--begin::Col-->
        <div class="d-none col-xxl-4">
          <!--begin::Forms widget 1-->
          <div class="card h-xl-100">
            <!--begin::Header-->
            <div class="card-header position-relative min-h-50px p-0 border-bottom-2">
              <!--begin::Nav-->
              <ul class="nav nav-pills nav-pills-custom d-flex position-relative w-100" role="tablist">
                <!--begin::Item-->
                <li class="nav-item mx-0 p-0 w-50" role="presentation">
                  <!--begin::Link-->
                  <a class="nav-link btn btn-color-muted active border-0 h-100 px-0" data-bs-toggle="pill" id="kt_forms_widget_1_tab_1" href="#kt_forms_widget_1_tab_content_1" aria-selected="true" role="tab">
                    <!--begin::Subtitle-->
                    <span class="nav-text fw-bold fs-4 mb-3">Buy</span>
                    <!--end::Subtitle-->
                    <!--begin::Bullet-->
                    <span class="bullet-custom position-absolute z-index-2 w-100 h-2px top-100 bottom-n100 bg-primary rounded"></span>
                    <!--end::Bullet-->
                  </a>
                  <!--end::Link-->
                </li>
                <!--end::Item-->
                <!--begin::Item-->
                <li class="nav-item mx-0 px-0 w-50" role="presentation">
                  <!--begin::Link-->
                  <a class="nav-link btn btn-color-muted border-0 h-100 px-0" data-bs-toggle="pill" id="kt_forms_widget_1_tab_2" href="#kt_forms_widget_1_tab_content_2" aria-selected="false" tabindex="-1" role="tab">
                    <!--begin::Subtitle-->
                    <span class="nav-text fw-bold fs-4 mb-3">Sell</span>
                    <!--end::Subtitle-->
                    <!--begin::Bullet-->
                    <span class="bullet-custom position-absolute z-index-2 w-100 h-2px top-100 bottom-n100 bg-primary rounded"></span>
                    <!--end::Bullet-->
                  </a>
                  <!--end::Link-->
                </li>
                <!--end::Item-->
              </ul>
              <!--end::Nav-->
            </div>
            <!--end::Header-->
            <!--begin::Body-->
            <div class="card-body">
              <!--begin::Tab Content-->
              <div class="tab-content">
                <!--begin::Tap pane-->
                <div class="tab-pane fade active show" id="kt_forms_widget_1_tab_content_1" role="tabpanel" aria-labelledby="#kt_forms_widget_1_tab_1">
                  <!--begin::Input group-->
                  <div class="form-floating border border-gray-300 rounded mb-7">
<!--                    <select class="form-select form-select-transparent select2-hidden-accessible" id="kt_forms_widget_1_select_1" data-select2-id="select2-data-kt_forms_widget_1_select_1" tabindex="-1" aria-hidden="true">-->
<!--                      <option></option>-->
<!--                      <option value="0" data-kt-select2-icon="assets/media/svg/coins/bitcoin.svg" selected="selected" data-select2-id="select2-data-128-m8ym">Bitcoin/BTC</option>-->
<!--                      <option value="1" data-kt-select2-icon="assets/media/svg/coins/ethereum.svg">Ethereum/ETH</option>-->
<!--                      <option value="2" data-kt-select2-icon="assets/media/svg/coins/filecoin.svg">Filecoin/FLE</option>-->
<!--                      <option value="3" data-kt-select2-icon="assets/media/svg/coins/chainlink.svg">Chainlink/CIN</option>-->
<!--                      <option value="4" data-kt-select2-icon="assets/media/svg/coins/binance.svg">Binance/BCN</option>-->
<!--                    </select><span class="select2 select2-container select2-container&#45;&#45;bootstrap5" dir="ltr" data-select2-id="select2-data-127-g75e" style="width: 100%;"><span class="selection"><span class="select2-selection select2-selection&#45;&#45;single form-select form-select-transparent" role="combobox" aria-haspopup="true" aria-expanded="false" tabindex="0" aria-disabled="false" aria-labelledby="select2-kt_forms_widget_1_select_1-container" aria-controls="select2-kt_forms_widget_1_select_1-container"><span class="select2-selection__rendered" id="select2-kt_forms_widget_1_select_1-container" role="textbox" aria-readonly="true" title="Bitcoin/BTC"><span><img src="assets/media/svg/coins/bitcoin.svg" class="rounded-circle h-20px me-2" alt="image">Bitcoin/BTC</span></span><span class="select2-selection__arrow" role="presentation"><b role="presentation"></b></span></span></span><span class="dropdown-wrapper" aria-hidden="true"></span></span>-->
                    <label for="floatingInputValue">Coin Name</label>
                  </div>
                  <!--end::Input group-->
                  <!--begin::Row-->
                  <div class="row mb-7">
                    <!--begin::Col-->
                    <div class="col-6">
                      <!--begin::Input group-->
                      <div class="form-floating">
                        <input type="email" class="form-control text-gray-800 fw-bold" placeholder="00.0" id="floatingInputValue" value="$230.00">
                        <label for="floatingInputValue">Amount(USD)</label>
                      </div>
                      <!--end::Input group-->
                    </div>
                    <!--end::Col-->
                    <!--begin::Col-->
                    <div class="col-6">
                      <!--begin::Input group-->
                      <div class="form-floating">
                        <input type="email" class="form-control text-gray-800 fw-bold" placeholder="00.0" id="floatingInputValue" value="$0,00000032">
                        <label for="floatingInputValue">Amount(BTC)</label>
                      </div>
                      <!--end::Input group-->
                    </div>
                    <!--end::Col-->
                  </div>
                  <!--end::Row-->
                  <!--begin::Action-->
                  <div class="d-flex align-items-end">
                    <a href="#" data-bs-toggle="modal" data-bs-target="#kt_modal_top_up_wallet" class="btn btn-primary fs-3 w-100">Make Payment</a>
                  </div>
                  <!--end::Action-->
                </div>
                <!--end::Tap pane-->
                <!--begin::Tap pane-->
                <div class="tab-pane fade" id="kt_forms_widget_1_tab_content_2" role="tabpanel" aria-labelledby="#kt_forms_widget_1_tab_2">
                  <!--begin::Input group-->
                  <div class="form-floating border rounded mb-7">
<!--                    <select class="form-select form-select-transparent select2-hidden-accessible" id="kt_forms_widget_1_select_2" data-select2-id="select2-data-kt_forms_widget_1_select_2" tabindex="-1" aria-hidden="true">-->
<!--                      <option></option>-->
<!--                      <option value="0" data-kt-select2-icon="assets/media/svg/coins/bitcoin.svg" selected="selected" data-select2-id="select2-data-130-5tb7">Bitcoin/BTC</option>-->
<!--                      <option value="1" data-kt-select2-icon="assets/media/svg/coins/ethereum.svg">Ethereum/ETH</option>-->
<!--                      <option value="2" data-kt-select2-icon="assets/media/svg/coins/filecoin.svg">Filecoin/FLE</option>-->
<!--                      <option value="3" data-kt-select2-icon="assets/media/svg/coins/chainlink.svg">Chainlink/CIN</option>-->
<!--                      <option value="4" data-kt-select2-icon="assets/media/svg/coins/binance.svg">Binance/BCN</option>-->
<!--                    </select><span class="select2 select2-container select2-container&#45;&#45;bootstrap5" dir="ltr" data-select2-id="select2-data-129-5lm0" style="width: 100%;"><span class="selection"><span class="select2-selection select2-selection&#45;&#45;single form-select form-select-transparent" role="combobox" aria-haspopup="true" aria-expanded="false" tabindex="0" aria-disabled="false" aria-labelledby="select2-kt_forms_widget_1_select_2-container" aria-controls="select2-kt_forms_widget_1_select_2-container"><span class="select2-selection__rendered" id="select2-kt_forms_widget_1_select_2-container" role="textbox" aria-readonly="true" title="Bitcoin/BTC"><span><img src="assets/media/svg/coins/bitcoin.svg" class="rounded-circle h-20px me-2" alt="image">Bitcoin/BTC</span></span><span class="select2-selection__arrow" role="presentation"><b role="presentation"></b></span></span></span><span class="dropdown-wrapper" aria-hidden="true"></span></span>-->
                    <label for="floatingInputValue">Coin Name</label>
                  </div>
                  <!--end::Input group-->
                  <!--begin::Row-->
                  <div class="row mb-7">
                    <!--begin::Col-->
                    <div class="col-6">
                      <!--begin::Input group-->
                      <div class="form-floating">
                        <input type="email" class="form-control text-gray-800 fw-bold" placeholder="00.0" id="floatingInputValue" value="$0,0000005">
                        <label for="floatingInputValue">Amount(BTC)</label>
                      </div>
                      <!--end::Input group-->
                    </div>
                    <!--end::Col-->
                    <!--begin::Col-->
                    <div class="col-6">
                      <!--begin::Input group-->
                      <div class="form-floating">
                        <input type="email" class="form-control text-gray-800 fw-bold" placeholder="00.0" id="floatingInputValue" value="$1230.00">
                        <label for="floatingInputValue">Amount(USD)</label>
                      </div>
                      <!--end::Input group-->
                    </div>
                    <!--end::Col-->
                  </div>
                  <!--end::Row-->
                  <!--begin::Action-->
                  <div class="d-flex align-items-end">
                    <a href="#" data-bs-toggle="modal" data-bs-target="#kt_modal_top_up_wallet" class="btn btn-primary fs-3 w-100">Place Offer</a>
                  </div>
                  <!--end::Action-->
                </div>
                <!--end::Tap pane-->
              </div>
              <!--end::Tab Content-->
            </div>
            <!--end: Card Body-->
          </div>
          <!--end::Forms widget 1-->
        </div>
        <!--end::Col-->
      </div>
      <!--end::Row-->
      <!--begin::Row-->
      <div class="row g-5 g-xl-10 mb-5 mb-xl-10 d-none">
        <!--begin::Col-->
        <div class="col-xxl-12">
          <!--begin::Chart widget 26-->
          <div class="card card-flush overflow-hidden h-xl-100">
            <!--begin::Header-->
            <div class="card-header pt-7 mb-2">
              <!--begin::Title-->
              <h3 class="card-title text-gray-800 fw-bold">Transaction History</h3>
              <!--end::Title-->
              <!--begin::Toolbar-->
              <div class="card-toolbar">
                <!--begin::Daterangepicker(defined in src/js/layout/app.js)-->
                <div data-kt-daterangepicker="true" data-kt-daterangepicker-opens="left" class="btn btn-sm btn-light d-flex align-items-center px-4" data-kt-initialized="1">
                  <!--begin::Display range-->
                  <div class="text-gray-600 fw-bold">23 авг. 2022 - 21 сент. 2022</div>
                  <!--end::Display range-->
                  <!--begin::Svg Icon | path: icons/duotune/general/gen014.svg-->
                  <span class="svg-icon svg-icon-1 ms-2 me-0">
															<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																<path opacity="0.3" d="M21 22H3C2.4 22 2 21.6 2 21V5C2 4.4 2.4 4 3 4H21C21.6 4 22 4.4 22 5V21C22 21.6 21.6 22 21 22Z" fill="currentColor"></path>
																<path d="M6 6C5.4 6 5 5.6 5 5V3C5 2.4 5.4 2 6 2C6.6 2 7 2.4 7 3V5C7 5.6 6.6 6 6 6ZM11 5V3C11 2.4 10.6 2 10 2C9.4 2 9 2.4 9 3V5C9 5.6 9.4 6 10 6C10.6 6 11 5.6 11 5ZM15 5V3C15 2.4 14.6 2 14 2C13.4 2 13 2.4 13 3V5C13 5.6 13.4 6 14 6C14.6 6 15 5.6 15 5ZM19 5V3C19 2.4 18.6 2 18 2C17.4 2 17 2.4 17 3V5C17 5.6 17.4 6 18 6C18.6 6 19 5.6 19 5Z" fill="currentColor"></path>
																<path d="M8.8 13.1C9.2 13.1 9.5 13 9.7 12.8C9.9 12.6 10.1 12.3 10.1 11.9C10.1 11.6 10 11.3 9.8 11.1C9.6 10.9 9.3 10.8 9 10.8C8.8 10.8 8.59999 10.8 8.39999 10.9C8.19999 11 8.1 11.1 8 11.2C7.9 11.3 7.8 11.4 7.7 11.6C7.6 11.8 7.5 11.9 7.5 12.1C7.5 12.2 7.4 12.2 7.3 12.3C7.2 12.4 7.09999 12.4 6.89999 12.4C6.69999 12.4 6.6 12.3 6.5 12.2C6.4 12.1 6.3 11.9 6.3 11.7C6.3 11.5 6.4 11.3 6.5 11.1C6.6 10.9 6.8 10.7 7 10.5C7.2 10.3 7.49999 10.1 7.89999 10C8.29999 9.90003 8.60001 9.80003 9.10001 9.80003C9.50001 9.80003 9.80001 9.90003 10.1 10C10.4 10.1 10.7 10.3 10.9 10.4C11.1 10.5 11.3 10.8 11.4 11.1C11.5 11.4 11.6 11.6 11.6 11.9C11.6 12.3 11.5 12.6 11.3 12.9C11.1 13.2 10.9 13.5 10.6 13.7C10.9 13.9 11.2 14.1 11.4 14.3C11.6 14.5 11.8 14.7 11.9 15C12 15.3 12.1 15.5 12.1 15.8C12.1 16.2 12 16.5 11.9 16.8C11.8 17.1 11.5 17.4 11.3 17.7C11.1 18 10.7 18.2 10.3 18.3C9.9 18.4 9.5 18.5 9 18.5C8.5 18.5 8.1 18.4 7.7 18.2C7.3 18 7 17.8 6.8 17.6C6.6 17.4 6.4 17.1 6.3 16.8C6.2 16.5 6.10001 16.3 6.10001 16.1C6.10001 15.9 6.2 15.7 6.3 15.6C6.4 15.5 6.6 15.4 6.8 15.4C6.9 15.4 7.00001 15.4 7.10001 15.5C7.20001 15.6 7.3 15.6 7.3 15.7C7.5 16.2 7.7 16.6 8 16.9C8.3 17.2 8.6 17.3 9 17.3C9.2 17.3 9.5 17.2 9.7 17.1C9.9 17 10.1 16.8 10.3 16.6C10.5 16.4 10.5 16.1 10.5 15.8C10.5 15.3 10.4 15 10.1 14.7C9.80001 14.4 9.50001 14.3 9.10001 14.3C9.00001 14.3 8.9 14.3 8.7 14.3C8.5 14.3 8.39999 14.3 8.39999 14.3C8.19999 14.3 7.99999 14.2 7.89999 14.1C7.79999 14 7.7 13.8 7.7 13.7C7.7 13.5 7.79999 13.4 7.89999 13.2C7.99999 13 8.2 13 8.5 13H8.8V13.1ZM15.3 17.5V12.2C14.3 13 13.6 13.3 13.3 13.3C13.1 13.3 13 13.2 12.9 13.1C12.8 13 12.7 12.8 12.7 12.6C12.7 12.4 12.8 12.3 12.9 12.2C13 12.1 13.2 12 13.6 11.8C14.1 11.6 14.5 11.3 14.7 11.1C14.9 10.9 15.2 10.6 15.5 10.3C15.8 10 15.9 9.80003 15.9 9.70003C15.9 9.60003 16.1 9.60004 16.3 9.60004C16.5 9.60004 16.7 9.70003 16.8 9.80003C16.9 9.90003 17 10.2 17 10.5V17.2C17 18 16.7 18.4 16.2 18.4C16 18.4 15.8 18.3 15.6 18.2C15.4 18.1 15.3 17.8 15.3 17.5Z" fill="currentColor"></path>
															</svg>
														</span>
                  <!--end::Svg Icon-->
                </div>
                <!--end::Daterangepicker-->
              </div>
              <!--end::Toolbar-->
            </div>
            <!--end::Header-->
            <!--begin::Card body-->
            <div class="card-body d-flex justify-content-between flex-column pt-0 pb-1 px-0">
              <!--begin::Info-->
              <div class="px-9 mb-5">
                <!--begin::Statistics-->
                <div class="d-flex align-items-center mb-2">
                  <!--begin::Currency-->
                  <span class="fs-4 fw-semibold text-gray-400 align-self-start me-1">$</span>
                  <!--end::Currency-->
                  <!--begin::Value-->
                  <span class="fs-2hx fw-bold text-gray-800 me-2 lh-1 ls-n2">12,706</span>
                  <!--end::Value-->
                  <!--begin::Label-->
                  <span class="badge badge-success fs-base">
														<!--begin::Svg Icon | path: icons/duotune/arrows/arr066.svg-->
														<span class="svg-icon svg-icon-5 svg-icon-white ms-n1">
															<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																<rect opacity="0.5" x="13" y="6" width="13" height="2" rx="1" transform="rotate(90 13 6)" fill="currentColor"></rect>
																<path d="M12.5657 8.56569L16.75 12.75C17.1642 13.1642 17.8358 13.1642 18.25 12.75C18.6642 12.3358 18.6642 11.6642 18.25 11.25L12.7071 5.70711C12.3166 5.31658 11.6834 5.31658 11.2929 5.70711L5.75 11.25C5.33579 11.6642 5.33579 12.3358 5.75 12.75C6.16421 13.1642 6.83579 13.1642 7.25 12.75L11.4343 8.56569C11.7467 8.25327 12.2533 8.25327 12.5657 8.56569Z" fill="currentColor"></path>
															</svg>
														</span>
                    <!--end::Svg Icon-->4.5%</span>
                  <!--end::Label-->
                </div>
                <!--end::Statistics-->
                <!--begin::Description-->
                <span class="fs-6 fw-semibold text-gray-400">Transactions in April</span>
                <!--end::Description-->
              </div>
              <!--end::Info-->
              <!--begin::Chart-->
              <div id="kt_charts_widget_26" class="min-h-auto ps-4 pe-6" data-kt-chart-info="Transactions" style="height: 300px; min-height: 315px;"><div id="apexchartsonsk7rwh" class="apexcharts-canvas apexchartsonsk7rwh apexcharts-theme-light" style="width: 616.5px; height: 300px;"><svg id="SvgjsSvg1254" width="616.5" height="300" xmlns="http://www.w3.org/2000/svg" version="1.1" xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:svgjs="http://svgjs.dev" class="apexcharts-svg apexcharts-zoomable" xmlns:data="ApexChartsNS" transform="translate(0, 0)" style="background: transparent;"><g id="SvgjsG1256" class="apexcharts-inner apexcharts-graphical" transform="translate(57.046875, 30)"><defs id="SvgjsDefs1255"><clipPath id="gridRectMaskonsk7rwh"><rect id="SvgjsRect1261" width="556.453125" height="224.82" x="-3.5" y="-1.5" rx="0" ry="0" opacity="1" stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fff"></rect></clipPath><clipPath id="forecastMaskonsk7rwh"></clipPath><clipPath id="nonForecastMaskonsk7rwh"></clipPath><clipPath id="gridRectMarkerMaskonsk7rwh"><rect id="SvgjsRect1262" width="553.453125" height="225.82" x="-2" y="-2" rx="0" ry="0" opacity="1" stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fff"></rect></clipPath><linearGradient id="SvgjsLinearGradient1267" x1="0" y1="0" x2="0" y2="1"><stop id="SvgjsStop1268" stop-opacity="0.4" stop-color="rgba(0,158,247,0.4)" offset="0"></stop><stop id="SvgjsStop1269" stop-opacity="0" stop-color="rgba(255,255,255,0)" offset="0.8"></stop><stop id="SvgjsStop1270" stop-opacity="0" stop-color="rgba(255,255,255,0)" offset="1"></stop></linearGradient></defs><g id="SvgjsG1273" class="apexcharts-xaxis" transform="translate(0, 0)"><g id="SvgjsG1274" class="apexcharts-xaxis-texts-g" transform="translate(0, -10)"><text id="SvgjsText1276" font-family="inherit" x="0" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1277"></tspan><title></title></text><text id="SvgjsText1279" font-family="inherit" x="30.525173611111107" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1280"></tspan><title></title></text><text id="SvgjsText1282" font-family="inherit" x="61.050347222222214" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1283"></tspan><title></title></text><text id="SvgjsText1285" font-family="inherit" x="91.57552083333333" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 92.71684265136719 239.32000732421875)"><tspan id="SvgjsTspan1286">Apr 04</tspan><title>Apr 04</title></text><text id="SvgjsText1288" font-family="inherit" x="122.10069444444444" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1289"></tspan><title></title></text><text id="SvgjsText1291" font-family="inherit" x="152.62586805555557" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1292"></tspan><title></title></text><text id="SvgjsText1294" font-family="inherit" x="183.15104166666669" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 184.21993255615234 239.32000732421875)"><tspan id="SvgjsTspan1295">Apr 07</tspan><title>Apr 07</title></text><text id="SvgjsText1297" font-family="inherit" x="213.6762152777778" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1298"></tspan><title></title></text><text id="SvgjsText1300" font-family="inherit" x="244.2013888888889" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1301"></tspan><title></title></text><text id="SvgjsText1303" font-family="inherit" x="274.7265625" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 275.9751281738281 239.32000732421875)"><tspan id="SvgjsTspan1304">Apr 10</tspan><title>Apr 10</title></text><text id="SvgjsText1306" font-family="inherit" x="305.2517361111111" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1307"></tspan><title></title></text><text id="SvgjsText1309" font-family="inherit" x="335.7769097222222" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1310"></tspan><title></title></text><text id="SvgjsText1312" font-family="inherit" x="366.30208333333326" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 367.4803466796875 239.32000732421875)"><tspan id="SvgjsTspan1313">Apr 13</tspan><title>Apr 13</title></text><text id="SvgjsText1315" font-family="inherit" x="396.82725694444434" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1316"></tspan><title></title></text><text id="SvgjsText1318" font-family="inherit" x="427.35243055555543" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1319"></tspan><title></title></text><text id="SvgjsText1321" font-family="inherit" x="457.8776041666665" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 458.8775939941406 239.32000732421875)"><tspan id="SvgjsTspan1322">Apr 18</tspan><title>Apr 18</title></text><text id="SvgjsText1324" font-family="inherit" x="488.4027777777776" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1325"></tspan><title></title></text><text id="SvgjsText1327" font-family="inherit" x="518.9279513888887" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1328"></tspan><title></title></text><text id="SvgjsText1330" font-family="inherit" x="549.4531249999998" y="244.82" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-xaxis-label " style="font-family: inherit;" transform="rotate(0 1 -1)"><tspan id="SvgjsTspan1331"></tspan><title></title></text></g></g><g id="SvgjsG1355" class="apexcharts-grid"><g id="SvgjsG1356" class="apexcharts-gridlines-horizontal"><line id="SvgjsLine1358" x1="0" y1="0" x2="549.453125" y2="0" stroke="#e4e6ef" stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line><line id="SvgjsLine1359" x1="0" y1="36.97" x2="549.453125" y2="36.97" stroke="#e4e6ef" stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line><line id="SvgjsLine1360" x1="0" y1="73.94" x2="549.453125" y2="73.94" stroke="#e4e6ef" stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line><line id="SvgjsLine1361" x1="0" y1="110.91" x2="549.453125" y2="110.91" stroke="#e4e6ef" stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line><line id="SvgjsLine1362" x1="0" y1="147.88" x2="549.453125" y2="147.88" stroke="#e4e6ef" stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line><line id="SvgjsLine1363" x1="0" y1="184.85" x2="549.453125" y2="184.85" stroke="#e4e6ef" stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line><line id="SvgjsLine1364" x1="0" y1="221.82" x2="549.453125" y2="221.82" stroke="#e4e6ef" stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line></g><g id="SvgjsG1357" class="apexcharts-gridlines-vertical"></g><line id="SvgjsLine1366" x1="0" y1="221.82" x2="549.453125" y2="221.82" stroke="transparent" stroke-dasharray="0" stroke-linecap="butt"></line><line id="SvgjsLine1365" x1="0" y1="1" x2="0" y2="221.82" stroke="transparent" stroke-dasharray="0" stroke-linecap="butt"></line></g><g id="SvgjsG1263" class="apexcharts-area-series apexcharts-plot-series"><g id="SvgjsG1264" class="apexcharts-series" seriesName="Transactions" data:longestSeries="true" rel="1" data:realIndex="0"><path id="SvgjsPath1271" d="M 0 221.82L 0 120.99272727272728C 10.683810763888888 120.99272727272728 19.84136284722222 120.99272727272728 30.525173611111107 120.99272727272728C 41.20898437499999 120.99272727272728 50.36653645833333 87.38363636363601 61.050347222222214 87.38363636363601C 71.7341579861111 87.38363636363601 80.89171006944444 87.38363636363601 91.57552083333333 87.38363636363601C 102.25933159722221 87.38363636363601 111.41688368055554 53.77454545454475 122.10069444444443 53.77454545454475C 132.78450520833331 53.77454545454475 141.94205729166666 53.77454545454475 152.62586805555554 53.77454545454475C 163.30967881944443 53.77454545454475 172.46723090277777 87.38363636363601 183.15104166666666 87.38363636363601C 193.83485243055554 87.38363636363601 202.99240451388889 87.38363636363601 213.67621527777777 87.38363636363601C 224.36002604166666 87.38363636363601 233.51757812499997 53.77454545454475 244.20138888888886 53.77454545454475C 254.88519965277777 53.77454545454475 264.0427517361111 53.77454545454475 274.7265625 53.77454545454475C 285.41037326388886 53.77454545454475 294.56792534722223 87.38363636363601 305.2517361111111 87.38363636363601C 315.93554687499994 87.38363636363601 325.0930989583333 87.38363636363601 335.7769097222222 87.38363636363601C 346.4607204861111 87.38363636363601 355.6182725694444 120.99272727272728 366.3020833333333 120.99272727272728C 376.9858940972222 120.99272727272728 386.14344618055554 120.99272727272728 396.8272569444444 120.99272727272728C 407.5110677083333 120.99272727272728 416.66861979166663 87.38363636363601 427.35243055555554 87.38363636363601C 438.0362413194444 87.38363636363601 447.19379340277777 87.38363636363601 457.87760416666663 87.38363636363601C 468.5614149305555 87.38363636363601 477.71896701388886 53.77454545454475 488.4027777777777 53.77454545454475C 499.0865885416666 53.77454545454475 508.24414062499994 53.77454545454475 518.9279513888888 53.77454545454475C 529.6117621527777 53.77454545454475 538.7693142361111 87.38363636363601 549.453125 87.38363636363601C 549.453125 87.38363636363601 549.453125 87.38363636363601 549.453125 221.82M 549.453125 87.38363636363601z" fill="url(#SvgjsLinearGradient1267)" fill-opacity="1" stroke-opacity="1" stroke-linecap="butt" stroke-width="0" stroke-dasharray="0" class="apexcharts-area" index="0" clip-path="url(#gridRectMaskonsk7rwh)" pathTo="M 0 221.82L 0 120.99272727272728C 10.683810763888888 120.99272727272728 19.84136284722222 120.99272727272728 30.525173611111107 120.99272727272728C 41.20898437499999 120.99272727272728 50.36653645833333 87.38363636363601 61.050347222222214 87.38363636363601C 71.7341579861111 87.38363636363601 80.89171006944444 87.38363636363601 91.57552083333333 87.38363636363601C 102.25933159722221 87.38363636363601 111.41688368055554 53.77454545454475 122.10069444444443 53.77454545454475C 132.78450520833331 53.77454545454475 141.94205729166666 53.77454545454475 152.62586805555554 53.77454545454475C 163.30967881944443 53.77454545454475 172.46723090277777 87.38363636363601 183.15104166666666 87.38363636363601C 193.83485243055554 87.38363636363601 202.99240451388889 87.38363636363601 213.67621527777777 87.38363636363601C 224.36002604166666 87.38363636363601 233.51757812499997 53.77454545454475 244.20138888888886 53.77454545454475C 254.88519965277777 53.77454545454475 264.0427517361111 53.77454545454475 274.7265625 53.77454545454475C 285.41037326388886 53.77454545454475 294.56792534722223 87.38363636363601 305.2517361111111 87.38363636363601C 315.93554687499994 87.38363636363601 325.0930989583333 87.38363636363601 335.7769097222222 87.38363636363601C 346.4607204861111 87.38363636363601 355.6182725694444 120.99272727272728 366.3020833333333 120.99272727272728C 376.9858940972222 120.99272727272728 386.14344618055554 120.99272727272728 396.8272569444444 120.99272727272728C 407.5110677083333 120.99272727272728 416.66861979166663 87.38363636363601 427.35243055555554 87.38363636363601C 438.0362413194444 87.38363636363601 447.19379340277777 87.38363636363601 457.87760416666663 87.38363636363601C 468.5614149305555 87.38363636363601 477.71896701388886 53.77454545454475 488.4027777777777 53.77454545454475C 499.0865885416666 53.77454545454475 508.24414062499994 53.77454545454475 518.9279513888888 53.77454545454475C 529.6117621527777 53.77454545454475 538.7693142361111 87.38363636363601 549.453125 87.38363636363601C 549.453125 87.38363636363601 549.453125 87.38363636363601 549.453125 221.82M 549.453125 87.38363636363601z" pathFrom="M -1 2440.020000000012L -1 2440.020000000012L 30.525173611111107 2440.020000000012L 61.050347222222214 2440.020000000012L 91.57552083333333 2440.020000000012L 122.10069444444443 2440.020000000012L 152.62586805555554 2440.020000000012L 183.15104166666666 2440.020000000012L 213.67621527777777 2440.020000000012L 244.20138888888886 2440.020000000012L 274.7265625 2440.020000000012L 305.2517361111111 2440.020000000012L 335.7769097222222 2440.020000000012L 366.3020833333333 2440.020000000012L 396.8272569444444 2440.020000000012L 427.35243055555554 2440.020000000012L 457.87760416666663 2440.020000000012L 488.4027777777777 2440.020000000012L 518.9279513888888 2440.020000000012L 549.453125 2440.020000000012"></path><path id="SvgjsPath1272" d="M 0 120.99272727272728C 10.683810763888888 120.99272727272728 19.84136284722222 120.99272727272728 30.525173611111107 120.99272727272728C 41.20898437499999 120.99272727272728 50.36653645833333 87.38363636363601 61.050347222222214 87.38363636363601C 71.7341579861111 87.38363636363601 80.89171006944444 87.38363636363601 91.57552083333333 87.38363636363601C 102.25933159722221 87.38363636363601 111.41688368055554 53.77454545454475 122.10069444444443 53.77454545454475C 132.78450520833331 53.77454545454475 141.94205729166666 53.77454545454475 152.62586805555554 53.77454545454475C 163.30967881944443 53.77454545454475 172.46723090277777 87.38363636363601 183.15104166666666 87.38363636363601C 193.83485243055554 87.38363636363601 202.99240451388889 87.38363636363601 213.67621527777777 87.38363636363601C 224.36002604166666 87.38363636363601 233.51757812499997 53.77454545454475 244.20138888888886 53.77454545454475C 254.88519965277777 53.77454545454475 264.0427517361111 53.77454545454475 274.7265625 53.77454545454475C 285.41037326388886 53.77454545454475 294.56792534722223 87.38363636363601 305.2517361111111 87.38363636363601C 315.93554687499994 87.38363636363601 325.0930989583333 87.38363636363601 335.7769097222222 87.38363636363601C 346.4607204861111 87.38363636363601 355.6182725694444 120.99272727272728 366.3020833333333 120.99272727272728C 376.9858940972222 120.99272727272728 386.14344618055554 120.99272727272728 396.8272569444444 120.99272727272728C 407.5110677083333 120.99272727272728 416.66861979166663 87.38363636363601 427.35243055555554 87.38363636363601C 438.0362413194444 87.38363636363601 447.19379340277777 87.38363636363601 457.87760416666663 87.38363636363601C 468.5614149305555 87.38363636363601 477.71896701388886 53.77454545454475 488.4027777777777 53.77454545454475C 499.0865885416666 53.77454545454475 508.24414062499994 53.77454545454475 518.9279513888888 53.77454545454475C 529.6117621527777 53.77454545454475 538.7693142361111 87.38363636363601 549.453125 87.38363636363601" fill="none" fill-opacity="1" stroke="#009ef7" stroke-opacity="1" stroke-linecap="butt" stroke-width="3" stroke-dasharray="0" class="apexcharts-area" index="0" clip-path="url(#gridRectMaskonsk7rwh)" pathTo="M 0 120.99272727272728C 10.683810763888888 120.99272727272728 19.84136284722222 120.99272727272728 30.525173611111107 120.99272727272728C 41.20898437499999 120.99272727272728 50.36653645833333 87.38363636363601 61.050347222222214 87.38363636363601C 71.7341579861111 87.38363636363601 80.89171006944444 87.38363636363601 91.57552083333333 87.38363636363601C 102.25933159722221 87.38363636363601 111.41688368055554 53.77454545454475 122.10069444444443 53.77454545454475C 132.78450520833331 53.77454545454475 141.94205729166666 53.77454545454475 152.62586805555554 53.77454545454475C 163.30967881944443 53.77454545454475 172.46723090277777 87.38363636363601 183.15104166666666 87.38363636363601C 193.83485243055554 87.38363636363601 202.99240451388889 87.38363636363601 213.67621527777777 87.38363636363601C 224.36002604166666 87.38363636363601 233.51757812499997 53.77454545454475 244.20138888888886 53.77454545454475C 254.88519965277777 53.77454545454475 264.0427517361111 53.77454545454475 274.7265625 53.77454545454475C 285.41037326388886 53.77454545454475 294.56792534722223 87.38363636363601 305.2517361111111 87.38363636363601C 315.93554687499994 87.38363636363601 325.0930989583333 87.38363636363601 335.7769097222222 87.38363636363601C 346.4607204861111 87.38363636363601 355.6182725694444 120.99272727272728 366.3020833333333 120.99272727272728C 376.9858940972222 120.99272727272728 386.14344618055554 120.99272727272728 396.8272569444444 120.99272727272728C 407.5110677083333 120.99272727272728 416.66861979166663 87.38363636363601 427.35243055555554 87.38363636363601C 438.0362413194444 87.38363636363601 447.19379340277777 87.38363636363601 457.87760416666663 87.38363636363601C 468.5614149305555 87.38363636363601 477.71896701388886 53.77454545454475 488.4027777777777 53.77454545454475C 499.0865885416666 53.77454545454475 508.24414062499994 53.77454545454475 518.9279513888888 53.77454545454475C 529.6117621527777 53.77454545454475 538.7693142361111 87.38363636363601 549.453125 87.38363636363601" pathFrom="M -1 2440.020000000012L -1 2440.020000000012L 30.525173611111107 2440.020000000012L 61.050347222222214 2440.020000000012L 91.57552083333333 2440.020000000012L 122.10069444444443 2440.020000000012L 152.62586805555554 2440.020000000012L 183.15104166666666 2440.020000000012L 213.67621527777777 2440.020000000012L 244.20138888888886 2440.020000000012L 274.7265625 2440.020000000012L 305.2517361111111 2440.020000000012L 335.7769097222222 2440.020000000012L 366.3020833333333 2440.020000000012L 396.8272569444444 2440.020000000012L 427.35243055555554 2440.020000000012L 457.87760416666663 2440.020000000012L 488.4027777777777 2440.020000000012L 518.9279513888888 2440.020000000012L 549.453125 2440.020000000012"></path><g id="SvgjsG1265" class="apexcharts-series-markers-wrap" data:realIndex="0"><g class="apexcharts-series-markers"><circle id="SvgjsCircle1374" r="0" cx="0" cy="0" class="apexcharts-marker wdph7bdk6k no-pointer-events" stroke="#009ef7" fill="#009ef7" fill-opacity="1" stroke-width="3" stroke-opacity="0.9" default-marker-size="0"></circle></g></g></g><g id="SvgjsG1266" class="apexcharts-datalabels" data:realIndex="0"></g></g><line id="SvgjsLine1368" x1="0" y1="0" x2="0" y2="221.82" stroke="#009ef7" stroke-dasharray="3" stroke-linecap="butt" class="apexcharts-xcrosshairs" x="0" y="0" width="1" height="221.82" fill="#b1b9c4" filter="none" fill-opacity="0.9" stroke-width="1"></line><line id="SvgjsLine1369" x1="0" y1="0" x2="549.453125" y2="0" stroke="#b6b6b6" stroke-dasharray="0" stroke-width="1" stroke-linecap="butt" class="apexcharts-ycrosshairs"></line><line id="SvgjsLine1370" x1="0" y1="0" x2="549.453125" y2="0" stroke-dasharray="0" stroke-width="0" stroke-linecap="butt" class="apexcharts-ycrosshairs-hidden"></line><g id="SvgjsG1371" class="apexcharts-yaxis-annotations"></g><g id="SvgjsG1372" class="apexcharts-xaxis-annotations"></g><g id="SvgjsG1373" class="apexcharts-point-annotations"></g><rect id="SvgjsRect1375" width="0" height="0" x="0" y="0" rx="0" ry="0" opacity="1" stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fefefe" class="apexcharts-zoom-rect"></rect><rect id="SvgjsRect1376" width="0" height="0" x="0" y="0" rx="0" ry="0" opacity="1" stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fefefe" class="apexcharts-selection-rect"></rect></g><g id="SvgjsG1332" class="apexcharts-yaxis" rel="0" transform="translate(27.046875, 0)"><g id="SvgjsG1333" class="apexcharts-yaxis-texts-g"><text id="SvgjsText1335" font-family="inherit" x="20" y="31.6" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-yaxis-label " style="font-family: inherit;"><tspan id="SvgjsTspan1336">$362</tspan><title>$362</title></text><text id="SvgjsText1338" font-family="inherit" x="20" y="68.57" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-yaxis-label " style="font-family: inherit;"><tspan id="SvgjsTspan1339">$357</tspan><title>$357</title></text><text id="SvgjsText1341" font-family="inherit" x="20" y="105.53999999999999" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-yaxis-label " style="font-family: inherit;"><tspan id="SvgjsTspan1342">$351</tspan><title>$351</title></text><text id="SvgjsText1344" font-family="inherit" x="20" y="142.51" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-yaxis-label " style="font-family: inherit;"><tspan id="SvgjsTspan1345">$346</tspan><title>$346</title></text><text id="SvgjsText1347" font-family="inherit" x="20" y="179.48" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-yaxis-label " style="font-family: inherit;"><tspan id="SvgjsTspan1348">$340</tspan><title>$340</title></text><text id="SvgjsText1350" font-family="inherit" x="20" y="216.45" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-yaxis-label " style="font-family: inherit;"><tspan id="SvgjsTspan1351">$335</tspan><title>$335</title></text><text id="SvgjsText1353" font-family="inherit" x="20" y="253.42" text-anchor="end" dominant-baseline="auto" font-size="12px" font-weight="400" fill="#a1a5b7" class="apexcharts-text apexcharts-yaxis-label " style="font-family: inherit;"><tspan id="SvgjsTspan1354">$330</tspan><title>$330</title></text></g></g><rect id="SvgjsRect1367" width="0" height="0" x="0" y="0" rx="0" ry="0" opacity="1" stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fefefe"></rect><g id="SvgjsG1257" class="apexcharts-annotations"></g></svg><div class="apexcharts-legend" style="max-height: 150px;"></div><div class="apexcharts-tooltip apexcharts-theme-light"><div class="apexcharts-tooltip-title" style="font-family: inherit; font-size: 12px;"></div><div class="apexcharts-tooltip-series-group" style="order: 1;"><span class="apexcharts-tooltip-marker" style="background-color: rgb(0, 158, 247);"></span><div class="apexcharts-tooltip-text" style="font-family: inherit; font-size: 12px;"><div class="apexcharts-tooltip-y-group"><span class="apexcharts-tooltip-text-y-label"></span><span class="apexcharts-tooltip-text-y-value"></span></div><div class="apexcharts-tooltip-goals-group"><span class="apexcharts-tooltip-text-goals-label"></span><span class="apexcharts-tooltip-text-goals-value"></span></div><div class="apexcharts-tooltip-z-group"><span class="apexcharts-tooltip-text-z-label"></span><span class="apexcharts-tooltip-text-z-value"></span></div></div></div></div><div class="apexcharts-xaxistooltip apexcharts-xaxistooltip-bottom apexcharts-theme-light"><div class="apexcharts-xaxistooltip-text" style="font-family: inherit; font-size: 12px;"></div></div><div class="apexcharts-yaxistooltip apexcharts-yaxistooltip-0 apexcharts-yaxistooltip-left apexcharts-theme-light"><div class="apexcharts-yaxistooltip-text"></div></div></div></div>
              <!--end::Chart-->
            </div>
            <!--end::Card body-->
          </div>
          <!--end::Chart widget 26-->
        </div>
        <!--end::Col-->
        <!--begin::Col-->
        <div class="col-xxl-4">
          <!--begin::Engage widget 1-->
          <div class="card h-md-100">
            <!--begin::Body-->
            <div class="card-body d-flex flex-column flex-center">
              <!--begin::Heading-->
              <div class="mb-2">
                <!--begin::Title-->
                <h1 class="fw-semibold text-gray-800 text-center lh-lg">Try out our
                  <br>new
                  <span class="fw-bolder">Invoice Manager</span></h1>
                <!--end::Title-->
                <!--begin::Illustration-->
<!--                <div class="py-10 text-center">-->
<!--                  <img src="assets/media/svg/illustrations/easy/2.svg" class="theme-light-show w-200px" alt="">-->
<!--                  <img src="assets/media/svg/illustrations/easy/2-dark.svg" class="theme-dark-show w-200px" alt="">-->
<!--                </div>-->
                <!--end::Illustration-->
              </div>
              <!--end::Heading-->
              <!--begin::Links-->
              <div class="text-center mb-1">
                <!--begin::Link-->
                <a class="btn btn-sm btn-primary me-2" data-bs-target="#kt_modal_create_account" data-bs-toggle="modal">Try Now</a>
                <!--end::Link-->
                <!--begin::Link-->
                <a class="btn btn-sm btn-light" href="../../demo14/dist/apps/ecommerce/sales/listing.html">Learn More</a>
                <!--end::Link-->
              </div>
              <!--end::Links-->
            </div>
            <!--end::Body-->
          </div>
          <!--end::Engage widget 1-->
        </div>
        <!--end::Col-->
      </div>
      <!--end::Row-->
      <!--begin::Row-->
      <div class="row g-5 g-xl-10 d-none">
        <!--begin::Col-->
        <div class="col-xxl-8">
          <!--begin::Tables widget 7-->
          <div class="card card-flush h-xl-100">
            <!--begin::Header-->
            <div class="card-header pt-7">
              <!--begin::Title-->
              <h4 class="card-title align-items-start flex-column">
                <span class="card-label fw-bold text-gray-800">Latest Activity</span>
                <span class="text-gray-400 mt-1 fw-semibold fs-7">Updated 37 minutes ago</span>
              </h4>
              <!--end::Title-->
              <!--begin::Toolbar-->
              <div class="card-toolbar">
                <!--begin::Daterangepicker(defined in src/js/layout/app.js)-->
                <div data-kt-daterangepicker="true" data-kt-daterangepicker-opens="left" class="btn btn-sm btn-light d-flex align-items-center px-4" data-kt-initialized="1">
                  <!--begin::Display range-->
                  <div class="text-gray-600 fw-bold">23 авг. 2022 - 21 сент. 2022</div>
                  <!--end::Display range-->
                  <!--begin::Svg Icon | path: icons/duotune/general/gen014.svg-->
                  <span class="svg-icon svg-icon-1 ms-2 me-0">
															<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																<path opacity="0.3" d="M21 22H3C2.4 22 2 21.6 2 21V5C2 4.4 2.4 4 3 4H21C21.6 4 22 4.4 22 5V21C22 21.6 21.6 22 21 22Z" fill="currentColor"></path>
																<path d="M6 6C5.4 6 5 5.6 5 5V3C5 2.4 5.4 2 6 2C6.6 2 7 2.4 7 3V5C7 5.6 6.6 6 6 6ZM11 5V3C11 2.4 10.6 2 10 2C9.4 2 9 2.4 9 3V5C9 5.6 9.4 6 10 6C10.6 6 11 5.6 11 5ZM15 5V3C15 2.4 14.6 2 14 2C13.4 2 13 2.4 13 3V5C13 5.6 13.4 6 14 6C14.6 6 15 5.6 15 5ZM19 5V3C19 2.4 18.6 2 18 2C17.4 2 17 2.4 17 3V5C17 5.6 17.4 6 18 6C18.6 6 19 5.6 19 5Z" fill="currentColor"></path>
																<path d="M8.8 13.1C9.2 13.1 9.5 13 9.7 12.8C9.9 12.6 10.1 12.3 10.1 11.9C10.1 11.6 10 11.3 9.8 11.1C9.6 10.9 9.3 10.8 9 10.8C8.8 10.8 8.59999 10.8 8.39999 10.9C8.19999 11 8.1 11.1 8 11.2C7.9 11.3 7.8 11.4 7.7 11.6C7.6 11.8 7.5 11.9 7.5 12.1C7.5 12.2 7.4 12.2 7.3 12.3C7.2 12.4 7.09999 12.4 6.89999 12.4C6.69999 12.4 6.6 12.3 6.5 12.2C6.4 12.1 6.3 11.9 6.3 11.7C6.3 11.5 6.4 11.3 6.5 11.1C6.6 10.9 6.8 10.7 7 10.5C7.2 10.3 7.49999 10.1 7.89999 10C8.29999 9.90003 8.60001 9.80003 9.10001 9.80003C9.50001 9.80003 9.80001 9.90003 10.1 10C10.4 10.1 10.7 10.3 10.9 10.4C11.1 10.5 11.3 10.8 11.4 11.1C11.5 11.4 11.6 11.6 11.6 11.9C11.6 12.3 11.5 12.6 11.3 12.9C11.1 13.2 10.9 13.5 10.6 13.7C10.9 13.9 11.2 14.1 11.4 14.3C11.6 14.5 11.8 14.7 11.9 15C12 15.3 12.1 15.5 12.1 15.8C12.1 16.2 12 16.5 11.9 16.8C11.8 17.1 11.5 17.4 11.3 17.7C11.1 18 10.7 18.2 10.3 18.3C9.9 18.4 9.5 18.5 9 18.5C8.5 18.5 8.1 18.4 7.7 18.2C7.3 18 7 17.8 6.8 17.6C6.6 17.4 6.4 17.1 6.3 16.8C6.2 16.5 6.10001 16.3 6.10001 16.1C6.10001 15.9 6.2 15.7 6.3 15.6C6.4 15.5 6.6 15.4 6.8 15.4C6.9 15.4 7.00001 15.4 7.10001 15.5C7.20001 15.6 7.3 15.6 7.3 15.7C7.5 16.2 7.7 16.6 8 16.9C8.3 17.2 8.6 17.3 9 17.3C9.2 17.3 9.5 17.2 9.7 17.1C9.9 17 10.1 16.8 10.3 16.6C10.5 16.4 10.5 16.1 10.5 15.8C10.5 15.3 10.4 15 10.1 14.7C9.80001 14.4 9.50001 14.3 9.10001 14.3C9.00001 14.3 8.9 14.3 8.7 14.3C8.5 14.3 8.39999 14.3 8.39999 14.3C8.19999 14.3 7.99999 14.2 7.89999 14.1C7.79999 14 7.7 13.8 7.7 13.7C7.7 13.5 7.79999 13.4 7.89999 13.2C7.99999 13 8.2 13 8.5 13H8.8V13.1ZM15.3 17.5V12.2C14.3 13 13.6 13.3 13.3 13.3C13.1 13.3 13 13.2 12.9 13.1C12.8 13 12.7 12.8 12.7 12.6C12.7 12.4 12.8 12.3 12.9 12.2C13 12.1 13.2 12 13.6 11.8C14.1 11.6 14.5 11.3 14.7 11.1C14.9 10.9 15.2 10.6 15.5 10.3C15.8 10 15.9 9.80003 15.9 9.70003C15.9 9.60003 16.1 9.60004 16.3 9.60004C16.5 9.60004 16.7 9.70003 16.8 9.80003C16.9 9.90003 17 10.2 17 10.5V17.2C17 18 16.7 18.4 16.2 18.4C16 18.4 15.8 18.3 15.6 18.2C15.4 18.1 15.3 17.8 15.3 17.5Z" fill="currentColor"></path>
															</svg>
														</span>
                  <!--end::Svg Icon-->
                </div>
                <!--end::Daterangepicker-->
              </div>
              <!--end::Toolbar-->
            </div>
            <!--end::Header-->
            <!--begin::Body-->
            <div class="card-body py-3">
              <!--begin::Table container-->
              <div class="table-responsive">
                <!--begin::Table-->
                <table class="table table-row-dashed align-middle gs-0 gy-4 my-0">
                  <!--begin::Table head-->
                  <thead>
                  <tr class="border-bottom-0">
                    <th class="p-0 w-50px"></th>
                    <th class="p-0 min-w-175px"></th>
                    <th class="p-0 min-w-175px"></th>
                    <th class="p-0 min-w-150px"></th>
                    <th class="p-0 min-w-150px"></th>
                    <th class="p-0 min-w-50px"></th>
                  </tr>
                  </thead>
                  <!--end::Table head-->
                  <!--begin::Table body-->
                  <tbody>
                  <tr>
                    <td>
                      <div class="symbol symbol-40px">
																		<span class="symbol-label bg-light-info">
																			<!--begin::Svg Icon | path: icons/duotune/abstract/abs025.svg-->
																			<span class="svg-icon svg-icon-2x svg-icon-info">
																				<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																					<path d="M16.925 3.90078V8.00077L12.025 10.8008V5.10078L15.525 3.10078C16.125 2.80078 16.925 3.20078 16.925 3.90078ZM2.525 13.5008L6.025 15.5008L10.925 12.7008L6.025 9.90078L2.525 11.9008C1.825 12.3008 1.825 13.2008 2.525 13.5008ZM18.025 19.7008V15.6008L13.125 12.8008V18.5008L16.625 20.5008C17.225 20.8008 18.025 20.4008 18.025 19.7008Z" fill="currentColor"></path>
																					<path opacity="0.3" d="M8.52499 3.10078L12.025 5.10078V10.8008L7.125 8.00077V3.90078C7.125 3.20078 7.92499 2.80078 8.52499 3.10078ZM7.42499 20.5008L10.925 18.5008V12.8008L6.02499 15.6008V19.7008C6.02499 20.4008 6.82499 20.8008 7.42499 20.5008ZM21.525 11.9008L18.025 9.90078L13.125 12.7008L18.025 15.5008L21.525 13.5008C22.225 13.2008 22.225 12.3008 21.525 11.9008Z" fill="currentColor"></path>
																				</svg>
																			</span>
                                      <!--end::Svg Icon-->
																		</span>
                      </div>
                    </td>
                    <td class="ps-0">
                      <a href="#" class="text-dark fw-bold text-hover-primary mb-1 fs-6">Insurance</a>
                      <span class="text-muted fw-semibold d-block fs-7">Personal Health</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">BTC Wallet</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Personal Account</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">23 Jan, 22</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Last Payment</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">-0.0024 BTC</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Balance</span>
                    </td>
                    <td class="text-end">
                      <a href="#" class="btn btn-sm btn-icon btn-bg-light btn-active-color-primary w-30px h-30px">
                        <!--begin::Svg Icon | path: icons/duotune/arrows/arr064.svg-->
                        <span class="svg-icon svg-icon-2">
																			<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																				<rect opacity="0.5" x="18" y="13" width="13" height="2" rx="1" transform="rotate(-180 18 13)" fill="currentColor"></rect>
																				<path d="M15.4343 12.5657L11.25 16.75C10.8358 17.1642 10.8358 17.8358 11.25 18.25C11.6642 18.6642 12.3358 18.6642 12.75 18.25L18.2929 12.7071C18.6834 12.3166 18.6834 11.6834 18.2929 11.2929L12.75 5.75C12.3358 5.33579 11.6642 5.33579 11.25 5.75C10.8358 6.16421 10.8358 6.83579 11.25 7.25L15.4343 11.4343C15.7467 11.7467 15.7467 12.2533 15.4343 12.5657Z" fill="currentColor"></path>
																			</svg>
																		</span>
                        <!--end::Svg Icon-->
                      </a>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <div class="symbol symbol-40px">
																		<span class="symbol-label bg-light-success">
																			<!--begin::Svg Icon | path: icons/duotune/medicine/med005.svg-->
																			<span class="svg-icon svg-icon-2x svg-icon-success">
																				<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																					<path opacity="0.3" d="M17.9061 13H11.2061C11.2061 12.4 10.8061 12 10.2061 12C9.60605 12 9.20605 12.4 9.20605 13H6.50606L9.20605 8.40002V4C8.60605 4 8.20605 3.6 8.20605 3C8.20605 2.4 8.60605 2 9.20605 2H15.2061C15.8061 2 16.2061 2.4 16.2061 3C16.2061 3.6 15.8061 4 15.2061 4V8.40002L17.9061 13ZM13.2061 9C12.6061 9 12.2061 9.4 12.2061 10C12.2061 10.6 12.6061 11 13.2061 11C13.8061 11 14.2061 10.6 14.2061 10C14.2061 9.4 13.8061 9 13.2061 9Z" fill="currentColor"></path>
																					<path d="M18.9061 22H5.40605C3.60605 22 2.40606 20 3.30606 18.4L6.40605 13H9.10605C9.10605 13.6 9.50605 14 10.106 14C10.706 14 11.106 13.6 11.106 13H17.8061L20.9061 18.4C21.9061 20 20.8061 22 18.9061 22ZM14.2061 15C13.1061 15 12.2061 15.9 12.2061 17C12.2061 18.1 13.1061 19 14.2061 19C15.3061 19 16.2061 18.1 16.2061 17C16.2061 15.9 15.3061 15 14.2061 15Z" fill="currentColor"></path>
																				</svg>
																			</span>
                                      <!--end::Svg Icon-->
																		</span>
                      </div>
                    </td>
                    <td class="ps-0">
                      <a href="#" class="text-dark fw-bold text-hover-primary mb-1 fs-6">Annette Black</a>
                      <span class="text-muted fw-semibold d-block fs-7">Zuid Area</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">ETH Wallet</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Personal Account</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">04 Feb, 22</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Last Payment</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">-0.346 ETH</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Balance</span>
                    </td>
                    <td class="text-end">
                      <a href="#" class="btn btn-sm btn-icon btn-bg-light btn-active-color-primary w-30px h-30px">
                        <!--begin::Svg Icon | path: icons/duotune/arrows/arr064.svg-->
                        <span class="svg-icon svg-icon-2">
																			<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																				<rect opacity="0.5" x="18" y="13" width="13" height="2" rx="1" transform="rotate(-180 18 13)" fill="currentColor"></rect>
																				<path d="M15.4343 12.5657L11.25 16.75C10.8358 17.1642 10.8358 17.8358 11.25 18.25C11.6642 18.6642 12.3358 18.6642 12.75 18.25L18.2929 12.7071C18.6834 12.3166 18.6834 11.6834 18.2929 11.2929L12.75 5.75C12.3358 5.33579 11.6642 5.33579 11.25 5.75C10.8358 6.16421 10.8358 6.83579 11.25 7.25L15.4343 11.4343C15.7467 11.7467 15.7467 12.2533 15.4343 12.5657Z" fill="currentColor"></path>
																			</svg>
																		</span>
                        <!--end::Svg Icon-->
                      </a>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <div class="symbol symbol-40px">
																		<span class="symbol-label bg-light-danger">
																			<!--begin::Svg Icon | path: icons/duotune/abstract/abs036.svg-->
																			<span class="svg-icon svg-icon-2x svg-icon-danger">
																				<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																					<path d="M22 12C22 12.2 22 12.5 22 12.7L19.5 10.2L16.9 12.8C16.9 12.5 17 12.3 17 12C17 9.5 15.2 7.50001 12.8 7.10001L10.2 4.5L12.7 2C17.9 2.4 22 6.7 22 12ZM11.2 16.9C8.80001 16.5 7 14.5 7 12C7 11.7 7.00001 11.5 7.10001 11.2L4.5 13.8L2 11.3C2 11.5 2 11.8 2 12C2 17.3 6.09999 21.6 11.3 22L13.8 19.5L11.2 16.9Z" fill="currentColor"></path>
																					<path opacity="0.3" d="M22 12.7C21.6 17.9 17.3 22 12 22C11.8 22 11.5 22 11.3 22L13.8 19.5L11.2 16.9C11.5 16.9 11.7 17 12 17C14.5 17 16.5 15.2 16.9 12.8L19.5 10.2L22 12.7ZM10.2 4.5L12.7 2C12.5 2 12.2 2 12 2C6.7 2 2.4 6.1 2 11.3L4.5 13.8L7.10001 11.2C7.50001 8.8 9.5 7 12 7C12.3 7 12.5 7.00001 12.8 7.10001L10.2 4.5Z" fill="currentColor"></path>
																				</svg>
																			</span>
                                      <!--end::Svg Icon-->
																		</span>
                      </div>
                    </td>
                    <td class="ps-0">
                      <a href="#" class="text-dark fw-bold text-hover-primary mb-1 fs-6">Esther Howard</a>
                      <span class="text-muted fw-semibold d-block fs-7">Zuid Area</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">BTC Wallet</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Personal Account</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">18 Feb, 22</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Last Payment</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">-0.00081 BTC</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Balance</span>
                    </td>
                    <td class="text-end">
                      <a href="#" class="btn btn-sm btn-icon btn-bg-light btn-active-color-primary w-30px h-30px">
                        <!--begin::Svg Icon | path: icons/duotune/arrows/arr064.svg-->
                        <span class="svg-icon svg-icon-2">
																			<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																				<rect opacity="0.5" x="18" y="13" width="13" height="2" rx="1" transform="rotate(-180 18 13)" fill="currentColor"></rect>
																				<path d="M15.4343 12.5657L11.25 16.75C10.8358 17.1642 10.8358 17.8358 11.25 18.25C11.6642 18.6642 12.3358 18.6642 12.75 18.25L18.2929 12.7071C18.6834 12.3166 18.6834 11.6834 18.2929 11.2929L12.75 5.75C12.3358 5.33579 11.6642 5.33579 11.25 5.75C10.8358 6.16421 10.8358 6.83579 11.25 7.25L15.4343 11.4343C15.7467 11.7467 15.7467 12.2533 15.4343 12.5657Z" fill="currentColor"></path>
																			</svg>
																		</span>
                        <!--end::Svg Icon-->
                      </a>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <div class="symbol symbol-40px">
																		<span class="symbol-label bg-light-primary">
																			<!--begin::Svg Icon | path: icons/duotune/abstract/abs020.svg-->
																			<span class="svg-icon svg-icon-2x svg-icon-primary">
																				<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																					<path d="M17.302 11.35L12.002 20.55H21.202C21.802 20.55 22.202 19.85 21.902 19.35L17.302 11.35Z" fill="currentColor"></path>
																					<path opacity="0.3" d="M12.002 20.55H2.802C2.202 20.55 1.80202 19.85 2.10202 19.35L6.70203 11.45L12.002 20.55ZM11.302 3.45L6.70203 11.35H17.302L12.702 3.45C12.402 2.85 11.602 2.85 11.302 3.45Z" fill="currentColor"></path>
																				</svg>
																			</span>
                                      <!--end::Svg Icon-->
																		</span>
                      </div>
                    </td>
                    <td class="ps-0">
                      <a href="#" class="text-dark fw-bold text-hover-primary mb-1 fs-6">Guy Hawkins</a>
                      <span class="text-muted fw-semibold d-block fs-7">Zuid Area</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">DOGE Wallet</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Personal Account</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">01 Apr, 22</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Last Payment</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">-456.34 DOGE</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Balance</span>
                    </td>
                    <td class="text-end">
                      <a href="#" class="btn btn-sm btn-icon btn-bg-light btn-active-color-primary w-30px h-30px">
                        <!--begin::Svg Icon | path: icons/duotune/arrows/arr064.svg-->
                        <span class="svg-icon svg-icon-2">
																			<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																				<rect opacity="0.5" x="18" y="13" width="13" height="2" rx="1" transform="rotate(-180 18 13)" fill="currentColor"></rect>
																				<path d="M15.4343 12.5657L11.25 16.75C10.8358 17.1642 10.8358 17.8358 11.25 18.25C11.6642 18.6642 12.3358 18.6642 12.75 18.25L18.2929 12.7071C18.6834 12.3166 18.6834 11.6834 18.2929 11.2929L12.75 5.75C12.3358 5.33579 11.6642 5.33579 11.25 5.75C10.8358 6.16421 10.8358 6.83579 11.25 7.25L15.4343 11.4343C15.7467 11.7467 15.7467 12.2533 15.4343 12.5657Z" fill="currentColor"></path>
																			</svg>
																		</span>
                        <!--end::Svg Icon-->
                      </a>
                    </td>
                  </tr>
                  <tr>
                    <td>
                      <div class="symbol symbol-40px">
																		<span class="symbol-label bg-light-warning">
																			<!--begin::Svg Icon | path: icons/duotune/technology/teh008.svg-->
																			<span class="svg-icon svg-icon-2x svg-icon-warning">
																				<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																					<path opacity="0.3" d="M11 6.5C11 9 9 11 6.5 11C4 11 2 9 2 6.5C2 4 4 2 6.5 2C9 2 11 4 11 6.5ZM17.5 2C15 2 13 4 13 6.5C13 9 15 11 17.5 11C20 11 22 9 22 6.5C22 4 20 2 17.5 2ZM6.5 13C4 13 2 15 2 17.5C2 20 4 22 6.5 22C9 22 11 20 11 17.5C11 15 9 13 6.5 13ZM17.5 13C15 13 13 15 13 17.5C13 20 15 22 17.5 22C20 22 22 20 22 17.5C22 15 20 13 17.5 13Z" fill="currentColor"></path>
																					<path d="M17.5 16C17.5 16 17.4 16 17.5 16L16.7 15.3C16.1 14.7 15.7 13.9 15.6 13.1C15.5 12.4 15.5 11.6 15.6 10.8C15.7 9.99999 16.1 9.19998 16.7 8.59998L17.4 7.90002H17.5C18.3 7.90002 19 7.20002 19 6.40002C19 5.60002 18.3 4.90002 17.5 4.90002C16.7 4.90002 16 5.60002 16 6.40002V6.5L15.3 7.20001C14.7 7.80001 13.9 8.19999 13.1 8.29999C12.4 8.39999 11.6 8.39999 10.8 8.29999C9.99999 8.19999 9.20001 7.80001 8.60001 7.20001L7.89999 6.5V6.40002C7.89999 5.60002 7.19999 4.90002 6.39999 4.90002C5.59999 4.90002 4.89999 5.60002 4.89999 6.40002C4.89999 7.20002 5.59999 7.90002 6.39999 7.90002H6.5L7.20001 8.59998C7.80001 9.19998 8.19999 9.99999 8.29999 10.8C8.39999 11.5 8.39999 12.3 8.29999 13.1C8.19999 13.9 7.80001 14.7 7.20001 15.3L6.5 16H6.39999C5.59999 16 4.89999 16.7 4.89999 17.5C4.89999 18.3 5.59999 19 6.39999 19C7.19999 19 7.89999 18.3 7.89999 17.5V17.4L8.60001 16.7C9.20001 16.1 9.99999 15.7 10.8 15.6C11.5 15.5 12.3 15.5 13.1 15.6C13.9 15.7 14.7 16.1 15.3 16.7L16 17.4V17.5C16 18.3 16.7 19 17.5 19C18.3 19 19 18.3 19 17.5C19 16.7 18.3 16 17.5 16Z" fill="currentColor"></path>
																				</svg>
																			</span>
                                      <!--end::Svg Icon-->
																		</span>
                      </div>
                    </td>
                    <td class="ps-0">
                      <a href="#" class="text-dark fw-bold text-hover-primary mb-1 fs-6">Marvin McKinney</a>
                      <span class="text-muted fw-semibold d-block fs-7">Zuid Area</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">BTC Wallet</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Personal Account</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">26 May, 22</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Last Payment</span>
                    </td>
                    <td>
                      <span class="text-dark fw-bold d-block fs-6">-0.000039 BTC</span>
                      <span class="text-gray-400 fw-semibold d-block fs-7">Balance</span>
                    </td>
                    <td class="text-end">
                      <a href="#" class="btn btn-sm btn-icon btn-bg-light btn-active-color-primary w-30px h-30px">
                        <!--begin::Svg Icon | path: icons/duotune/arrows/arr064.svg-->
                        <span class="svg-icon svg-icon-2">
																			<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																				<rect opacity="0.5" x="18" y="13" width="13" height="2" rx="1" transform="rotate(-180 18 13)" fill="currentColor"></rect>
																				<path d="M15.4343 12.5657L11.25 16.75C10.8358 17.1642 10.8358 17.8358 11.25 18.25C11.6642 18.6642 12.3358 18.6642 12.75 18.25L18.2929 12.7071C18.6834 12.3166 18.6834 11.6834 18.2929 11.2929L12.75 5.75C12.3358 5.33579 11.6642 5.33579 11.25 5.75C10.8358 6.16421 10.8358 6.83579 11.25 7.25L15.4343 11.4343C15.7467 11.7467 15.7467 12.2533 15.4343 12.5657Z" fill="currentColor"></path>
																			</svg>
																		</span>
                        <!--end::Svg Icon-->
                      </a>
                    </td>
                  </tr>
                  </tbody>
                  <!--end::Table body-->
                </table>
                <!--end::Table-->
              </div>
              <!--end::Table container-->
            </div>
            <!--begin::Body-->
          </div>
          <!--end::Tables widget 7-->
        </div>
        <!--end::Col-->
        <!--begin::Col-->
        <div class="col-xxl-4">
          <!--begin::List widget 22-->
          <div class="card card-flush h-xl-100">
            <!--begin::Header-->
            <div class="card-header pt-5">
              <!--begin::Title-->
              <h3 class="card-title align-items-start flex-column">
                <span class="card-label fw-bold text-gray-800">Mining Status</span>
                <span class="text-gray-400 mt-1 fw-semibold fs-6">8k social visitors</span>
              </h3>
              <!--end::Title-->
              <!--begin::Toolbar-->
              <div class="card-toolbar">
                <a href="#" class="btn btn-sm btn-light">All Courses</a>
              </div>
              <!--end::Toolbar-->
            </div>
            <!--end::Header-->
            <!--begin::Body-->
            <div class="card-body">
              <!--begin::Item-->
              <div class="d-flex flex-stack">
                <!--begin::Section-->
                <div class="d-flex align-items-center me-5">
                  <!--begin::Symbol-->
                  <div class="symbol symbol-40px me-3">
															<span class="symbol-label bg-light-info">
																<!--begin::Svg Icon | path: icons/duotune/abstract/abs025.svg-->
																<span class="svg-icon svg-icon-2x svg-icon-info">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<path d="M16.925 3.90078V8.00077L12.025 10.8008V5.10078L15.525 3.10078C16.125 2.80078 16.925 3.20078 16.925 3.90078ZM2.525 13.5008L6.025 15.5008L10.925 12.7008L6.025 9.90078L2.525 11.9008C1.825 12.3008 1.825 13.2008 2.525 13.5008ZM18.025 19.7008V15.6008L13.125 12.8008V18.5008L16.625 20.5008C17.225 20.8008 18.025 20.4008 18.025 19.7008Z" fill="currentColor"></path>
																		<path opacity="0.3" d="M8.52499 3.10078L12.025 5.10078V10.8008L7.125 8.00077V3.90078C7.125 3.20078 7.92499 2.80078 8.52499 3.10078ZM7.42499 20.5008L10.925 18.5008V12.8008L6.02499 15.6008V19.7008C6.02499 20.4008 6.82499 20.8008 7.42499 20.5008ZM21.525 11.9008L18.025 9.90078L13.125 12.7008L18.025 15.5008L21.525 13.5008C22.225 13.2008 22.225 12.3008 21.525 11.9008Z" fill="currentColor"></path>
																	</svg>
																</span>
                                <!--end::Svg Icon-->
															</span>
                  </div>
                  <!--end::Symbol-->
                  <!--begin::Content-->
                  <div class="me-5">
                    <!--begin::Title-->
                    <a href="#" class="text-gray-800 fw-bold text-hover-primary fs-6">GPUs mining</a>
                    <!--end::Title-->
                    <!--begin::Desc-->
                    <span class="fw-semibold fs-7 d-block text-start text-success ps-0">Running</span>
                    <!--end::Desc-->
                  </div>
                  <!--end::Content-->
                </div>
                <!--end::Section-->
                <!--begin::Wrapper-->
                <div class="d-flex align-items-center">
                  <!--begin::Info-->
                  <div class="d-flex flex-center">
                    <!--begin::Action-->
                    <div class="form-check form-switch form-check-custom form-check-solid">
                      <input class="form-check-input h-20px w-30px" type="checkbox" value="" id="flexSwitchChecked" checked="checked">
                    </div>
                    <!--end::Action-->
                  </div>
                  <!--end::Info-->
                </div>
                <!--end::Wrapper-->
              </div>
              <!--end::Item-->
              <!--begin::Separator-->
              <div class="separator separator-dashed my-4"></div>
              <!--end::Separator-->
              <!--begin::Item-->
              <div class="d-flex flex-stack">
                <!--begin::Section-->
                <div class="d-flex align-items-center me-5">
                  <!--begin::Symbol-->
                  <div class="symbol symbol-40px me-3">
															<span class="symbol-label bg-light-success">
																<!--begin::Svg Icon | path: icons/duotune/medicine/med005.svg-->
																<span class="svg-icon svg-icon-2x svg-icon-success">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<path opacity="0.3" d="M17.9061 13H11.2061C11.2061 12.4 10.8061 12 10.2061 12C9.60605 12 9.20605 12.4 9.20605 13H6.50606L9.20605 8.40002V4C8.60605 4 8.20605 3.6 8.20605 3C8.20605 2.4 8.60605 2 9.20605 2H15.2061C15.8061 2 16.2061 2.4 16.2061 3C16.2061 3.6 15.8061 4 15.2061 4V8.40002L17.9061 13ZM13.2061 9C12.6061 9 12.2061 9.4 12.2061 10C12.2061 10.6 12.6061 11 13.2061 11C13.8061 11 14.2061 10.6 14.2061 10C14.2061 9.4 13.8061 9 13.2061 9Z" fill="currentColor"></path>
																		<path d="M18.9061 22H5.40605C3.60605 22 2.40606 20 3.30606 18.4L6.40605 13H9.10605C9.10605 13.6 9.50605 14 10.106 14C10.706 14 11.106 13.6 11.106 13H17.8061L20.9061 18.4C21.9061 20 20.8061 22 18.9061 22ZM14.2061 15C13.1061 15 12.2061 15.9 12.2061 17C12.2061 18.1 13.1061 19 14.2061 19C15.3061 19 16.2061 18.1 16.2061 17C16.2061 15.9 15.3061 15 14.2061 15Z" fill="currentColor"></path>
																	</svg>
																</span>
                                <!--end::Svg Icon-->
															</span>
                  </div>
                  <!--end::Symbol-->
                  <!--begin::Content-->
                  <div class="me-5">
                    <!--begin::Title-->
                    <a href="#" class="text-gray-800 fw-bold text-hover-primary fs-6">GPUs mining</a>
                    <!--end::Title-->
                    <!--begin::Desc-->
                    <span class="fw-semibold fs-7 d-block text-start text-success ps-0">Running</span>
                    <!--end::Desc-->
                  </div>
                  <!--end::Content-->
                </div>
                <!--end::Section-->
                <!--begin::Wrapper-->
                <div class="d-flex align-items-center">
                  <!--begin::Info-->
                  <div class="d-flex flex-center">
                    <!--begin::Action-->
                    <div class="form-check form-switch form-check-custom form-check-solid">
                      <input class="form-check-input h-20px w-30px" type="checkbox" value="" id="flexSwitchChecked" checked="checked">
                    </div>
                    <!--end::Action-->
                  </div>
                  <!--end::Info-->
                </div>
                <!--end::Wrapper-->
              </div>
              <!--end::Item-->
              <!--begin::Separator-->
              <div class="separator separator-dashed my-4"></div>
              <!--end::Separator-->
              <!--begin::Item-->
              <div class="d-flex flex-stack">
                <!--begin::Section-->
                <div class="d-flex align-items-center me-5">
                  <!--begin::Symbol-->
                  <div class="symbol symbol-40px me-3">
															<span class="symbol-label bg-light-danger">
																<!--begin::Svg Icon | path: icons/duotune/abstract/abs036.svg-->
																<span class="svg-icon svg-icon-2x svg-icon-danger">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<path d="M22 12C22 12.2 22 12.5 22 12.7L19.5 10.2L16.9 12.8C16.9 12.5 17 12.3 17 12C17 9.5 15.2 7.50001 12.8 7.10001L10.2 4.5L12.7 2C17.9 2.4 22 6.7 22 12ZM11.2 16.9C8.80001 16.5 7 14.5 7 12C7 11.7 7.00001 11.5 7.10001 11.2L4.5 13.8L2 11.3C2 11.5 2 11.8 2 12C2 17.3 6.09999 21.6 11.3 22L13.8 19.5L11.2 16.9Z" fill="currentColor"></path>
																		<path opacity="0.3" d="M22 12.7C21.6 17.9 17.3 22 12 22C11.8 22 11.5 22 11.3 22L13.8 19.5L11.2 16.9C11.5 16.9 11.7 17 12 17C14.5 17 16.5 15.2 16.9 12.8L19.5 10.2L22 12.7ZM10.2 4.5L12.7 2C12.5 2 12.2 2 12 2C6.7 2 2.4 6.1 2 11.3L4.5 13.8L7.10001 11.2C7.50001 8.8 9.5 7 12 7C12.3 7 12.5 7.00001 12.8 7.10001L10.2 4.5Z" fill="currentColor"></path>
																	</svg>
																</span>
                                <!--end::Svg Icon-->
															</span>
                  </div>
                  <!--end::Symbol-->
                  <!--begin::Content-->
                  <div class="me-5">
                    <!--begin::Title-->
                    <a href="#" class="text-gray-800 fw-bold text-hover-primary fs-6">Est. daily USD</a>
                    <!--end::Title-->
                    <!--begin::Desc-->
                    <span class="fw-semibold fs-7 d-block text-start text-gray-400 ps-0">$48.02</span>
                    <!--end::Desc-->
                  </div>
                  <!--end::Content-->
                </div>
                <!--end::Section-->
                <!--begin::Wrapper-->
                <div class="d-flex align-items-center">
                  <!--begin::Info-->
                  <div class="d-flex flex-center">
                    <!--begin::Action-->
                    <a href="#" class="btn btn-sm btn-icon btn-bg-light btn-active-color-primary w-25px h-25px">
                      <!--begin::Svg Icon | path: icons/duotune/arrows/arr001.svg-->
                      <span class="svg-icon svg-icon-5 svg-icon-gray-700">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<path d="M14.4 11H3C2.4 11 2 11.4 2 12C2 12.6 2.4 13 3 13H14.4V11Z" fill="currentColor"></path>
																		<path opacity="0.3" d="M14.4 20V4L21.7 11.3C22.1 11.7 22.1 12.3 21.7 12.7L14.4 20Z" fill="currentColor"></path>
																	</svg>
																</span>
                      <!--end::Svg Icon-->
                    </a>
                    <!--end::Action-->
                  </div>
                  <!--end::Info-->
                </div>
                <!--end::Wrapper-->
              </div>
              <!--end::Item-->
              <!--begin::Separator-->
              <div class="separator separator-dashed my-4"></div>
              <!--end::Separator-->
              <!--begin::Item-->
              <div class="d-flex flex-stack">
                <!--begin::Section-->
                <div class="d-flex align-items-center me-5">
                  <!--begin::Symbol-->
                  <div class="symbol symbol-40px me-3">
															<span class="symbol-label bg-light-primary">
																<!--begin::Svg Icon | path: icons/duotune/abstract/abs020.svg-->
																<span class="svg-icon svg-icon-2x svg-icon-primary">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<path d="M17.302 11.35L12.002 20.55H21.202C21.802 20.55 22.202 19.85 21.902 19.35L17.302 11.35Z" fill="currentColor"></path>
																		<path opacity="0.3" d="M12.002 20.55H2.802C2.202 20.55 1.80202 19.85 2.10202 19.35L6.70203 11.45L12.002 20.55ZM11.302 3.45L6.70203 11.35H17.302L12.702 3.45C12.402 2.85 11.602 2.85 11.302 3.45Z" fill="currentColor"></path>
																	</svg>
																</span>
                                <!--end::Svg Icon-->
															</span>
                  </div>
                  <!--end::Symbol-->
                  <!--begin::Content-->
                  <div class="me-5">
                    <!--begin::Title-->
                    <a href="#" class="text-gray-800 fw-bold text-hover-primary fs-6">Team Members</a>
                    <!--end::Title-->
                    <!--begin::Desc-->
                    <span class="fw-semibold fs-7 d-block text-start text-gray-400 ps-0">6</span>
                    <!--end::Desc-->
                  </div>
                  <!--end::Content-->
                </div>
                <!--end::Section-->
                <!--begin::Wrapper-->
                <div class="d-flex align-items-center">
                  <!--begin::Info-->
                  <div class="d-flex flex-center">
                    <!--begin::Action-->
                    <a href="#" class="btn btn-sm btn-icon btn-bg-light btn-active-color-primary w-25px h-25px">
                      <!--begin::Svg Icon | path: icons/duotune/arrows/arr001.svg-->
                      <span class="svg-icon svg-icon-5 svg-icon-gray-700">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
																		<path d="M14.4 11H3C2.4 11 2 11.4 2 12C2 12.6 2.4 13 3 13H14.4V11Z" fill="currentColor"></path>
																		<path opacity="0.3" d="M14.4 20V4L21.7 11.3C22.1 11.7 22.1 12.3 21.7 12.7L14.4 20Z" fill="currentColor"></path>
																	</svg>
																</span>
                      <!--end::Svg Icon-->
                    </a>
                    <!--end::Action-->
                  </div>
                  <!--end::Info-->
                </div>
                <!--end::Wrapper-->
              </div>
              <!--end::Item-->
            </div>
            <!--end::Body-->
            <!--begin::Footer-->
            <div class="card-footer mx-auto pt-0">
              <!--begin::Actions-->
              <a href="#" class="btn btn-primary btn-sm me-3" data-bs-toggle="modal" data-bs-target="#kt_modal_view_users">Add PC</a>
              <a href="#" class="btn btn-light btn-sm" data-bs-toggle="modal" data-bs-target="#kt_modal_users_search">Buy GPU</a>
              <!--end::Actions-->
            </div>
            <!--end::Footer-->
          </div>
          <!--end::List widget 22-->
        </div>
        <!--end::Col-->
      </div>
      <!--end::Row-->
    </div>
    <!--end::Post-->
    <!--begin::Footer-->
    <div class="footer py-4 d-flex flex-column flex-md-row flex-stack d-none" id="kt_footer">
      <!--begin::Copyright-->
      <div class="text-dark order-2 order-md-1">
        <span class="text-muted fw-semibold me-1">2022©</span>
        <a href="https://keenthemes.com" target="_blank" class="text-gray-800 text-hover-primary">Keenthemes</a>
      </div>
      <!--end::Copyright-->
      <!--begin::Menu-->
      <ul class="menu menu-gray-600 menu-hover-primary fw-semibold order-1">
        <li class="menu-item">
          <a href="https://keenthemes.com" target="_blank" class="menu-link px-2">About</a>
        </li>
        <li class="menu-item">
          <a href="https://devs.keenthemes.com" target="_blank" class="menu-link px-2">Support</a>
        </li>
        <li class="menu-item">
          <a href="https://1.envato.market/EA4JP" target="_blank" class="menu-link px-2">Purchase</a>
        </li>
      </ul>
      <!--end::Menu-->
    </div>
    <!--end::Footer-->
  </div>
</template>

<script>
import {mapGetters} from "vuex";

export default {
  name: "userAccounts",
  computed: {
    ...mapGetters({
      userinfo: "users/getUserInfo",
      userInfoLoaded: "users/getUserInfoLoaded",
      cards: "users/getCards",
      cardsLoaded: "users/getCardsLoaded",
      loggedIn: "auth/isLoggedIn"
    }),
  },
  methods: {
    getDataInfo() {
      if (this.loggedIn && !this.userInfoLoaded) {
        this.$store.dispatch("users/getUserInfoMethod");
      }
      if (this.loggedIn && !this.cardsLoaded) {
        this.$store.dispatch("users/getCardsMethod");
      }
      return true
    },
    totalBalance() {
      if (this.userinfo && this.cards.length > 0) {
        let cards = Object.values(this.cards)
        let card = cards.find(item => item.userId === this.userinfo.ID)
        return this.userinfo.USDTBalance + this.userinfo.USDBalance + this.userinfo.BTCBalance + this.userinfo.EURBalance + card.balance + "$";
      }
      return "You have no issued cards";
    },
  }
}
</script>

<style scoped>

</style>