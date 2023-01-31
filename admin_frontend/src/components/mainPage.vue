<template>
  <div v-if="loggedIn">
    <div v-if="getData()" class="d-flex flex-column flex-column-fluid container-fluid">
      <!--begin::Toolbar-->
      <div class="toolbar mb-5 mb-lg-7" id="kt_toolbar">
        <!--begin::Page title-->
        <div class="page-title d-flex flex-column me-3">
          <!--begin::Title-->
          <h1 class="d-flex text-dark fw-bold my-1 fs-3">Store Analytics</h1>
          <!--end::Title-->
        </div>
        <!--end::Page title-->
        <!--begin::Actions-->
        <div class="d-flex align-items-center py-2 py-md-1">
          <!--begin::Wrapper-->
          <div class="me-3" data-select2-id="select2-data-121-xm9k">
            <!--begin::Menu-->
            <a href="#" class="btn btn-light fw-bold" data-kt-menu-trigger="click" data-kt-menu-placement="bottom-end">
              <!--begin::Svg Icon | path: icons/duotune/general/gen031.svg-->
              <span class="svg-icon svg-icon-5 svg-icon-gray-500 me-1">
											<svg width="24" height="24" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
												<path
                            d="M19.0759 3H4.72777C3.95892 3 3.47768 3.83148 3.86067 4.49814L8.56967 12.6949C9.17923 13.7559 9.5 14.9582 9.5 16.1819V19.5072C9.5 20.2189 10.2223 20.7028 10.8805 20.432L13.8805 19.1977C14.2553 19.0435 14.5 18.6783 14.5 18.273V13.8372C14.5 12.8089 14.8171 11.8056 15.408 10.964L19.8943 4.57465C20.3596 3.912 19.8856 3 19.0759 3Z"
                            fill="currentColor"></path>
											</svg>
										</span>
              <!--end::Svg Icon-->Filter</a>
            <!--begin::Menu 1-->
            <div class="menu menu-sub menu-sub-dropdown w-250px w-md-300px" data-kt-menu="true"
                 id="kt_menu_62cfb07ef1e97" style="">
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
                <div class="mb-10" data-select2-id="select2-data-120-3829">
                  <!--begin::Label-->
                  <label class="form-label fw-semibold">Status:</label>
                  <!--end::Label-->
                  <!--begin::Input-->
                  <div data-select2-id="select2-data-119-pi3o">
                    <select class="form-select form-select-solid select2-hidden-accessible" data-kt-select2="true"
                            data-placeholder="Select option" data-dropdown-parent="#kt_menu_62cfb07ef1e97"
                            data-allow-clear="true" data-select2-id="select2-data-7-nry4" tabindex="-1"
                            aria-hidden="true" data-kt-initialized="1">
                      <option data-select2-id="select2-data-9-oyvq"></option>
                      <option value="1" data-select2-id="select2-data-123-3vyi">Approved</option>
                      <option value="2" data-select2-id="select2-data-124-gpux">Pending</option>
                      <option value="2" data-select2-id="select2-data-125-ky8i">In Process</option>
                      <option value="2" data-select2-id="select2-data-126-y9ot">Rejected</option>
                    </select><span
                      class="select2 select2-container select2-container--bootstrap5 select2-container--below" dir="ltr"
                      data-select2-id="select2-data-8-5b26" style="width: 100%;"><span class="selection"><span
                      class="select2-selection select2-selection--single form-select form-select-solid select2-selection--clearable"
                      role="combobox" aria-haspopup="true" aria-expanded="false" tabindex="0" aria-disabled="false"
                      aria-labelledby="select2-fmlq-container" aria-controls="select2-fmlq-container"><button
                      type="button" class="select2-selection__clear" tabindex="-1" title="Remove all items"
                      aria-label="Remove all items" aria-describedby="select2-fmlq-container"
                      data-select2-id="select2-data-127-4ysq"><span aria-hidden="true">×</span></button><span
                      class="select2-selection__rendered" id="select2-fmlq-container" role="textbox"
                      aria-readonly="true" title="Pending">Pending</span><span class="select2-selection__arrow"
                                                                               role="presentation"><b
                      role="presentation"></b></span></span></span><span class="dropdown-wrapper"
                                                                         aria-hidden="true"></span></span>
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
                  <button type="reset" class="btn btn-sm btn-light btn-active-light-primary me-2"
                          data-kt-menu-dismiss="true">Reset
                  </button>
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
          <a href="#" class="btn btn-dark fw-bold" data-bs-toggle="modal" data-bs-target="#kt_modal_create_app"
             id="kt_toolbar_primary_button">Create</a>
          <!--end::Button-->
        </div>
        <!--end::Actions-->
      </div>
      <!--end::Toolbar-->
      <!--begin::Post-->
      <div class="content flex-column-fluid" id="kt_content">
        <!--begin::Row-->
        <div class="row g-5 g-xl-10">
          <!--begin::Col-->
          <div class="col-xxl-6 mb-md-5 mb-xl-10">
            <!--begin::Row-->
            <div class="row g-5 g-xl-10">
              <!--begin::Col-->
              <div class="col-md-6 col-xl-6 mb-xxl-10">
                <!--begin::Card widget 8-->
                <div class="card overflow-hidden h-md-50 mb-5 mb-xl-10">
                  <!--begin::Card body-->
                  <div class="card-body d-flex justify-content-between flex-column px-0 pb-0">
                    <!--begin::Statistics-->
                    <div class="mb-4 px-9">
                      <!--begin::Info-->
                      <div class="d-flex align-items-center mb-2">
                        <!--begin::Value-->
                        <span class="fs-2hx fw-bold text-gray-800 me-2 lh-1">{{this.users.length}}</span>
                        <!--end::Value-->
                        <!--begin::Label-->
                        <span class="badge badge-light-success fs-base">
																<!--begin::Svg Icon | path: icons/duotune/arrows/arr066.svg-->
																<span class="svg-icon svg-icon-5 svg-icon-success ms-n1">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                       xmlns="http://www.w3.org/2000/svg">
																		<rect opacity="0.5" x="13" y="6" width="13" height="2" rx="1"
                                          transform="rotate(90 13 6)" fill="currentColor"></rect>
																		<path
                                        d="M12.5657 8.56569L16.75 12.75C17.1642 13.1642 17.8358 13.1642 18.25 12.75C18.6642 12.3358 18.6642 11.6642 18.25 11.25L12.7071 5.70711C12.3166 5.31658 11.6834 5.31658 11.2929 5.70711L5.75 11.25C5.33579 11.6642 5.33579 12.3358 5.75 12.75C6.16421 13.1642 6.83579 13.1642 7.25 12.75L11.4343 8.56569C11.7467 8.25327 12.2533 8.25327 12.5657 8.56569Z"
                                        fill="currentColor"></path>
																	</svg>
																</span>
                          <!--end::Svg Icon-->{{dataAdded(this.users)}}</span>
                        <!--end::Label-->
                      </div>
                      <!--end::Info-->
                      <!--begin::Description-->
                      <span class="fs-6 fw-semibold text-gray-400">Total Users</span>
                      <!--end::Description-->
                    </div>
                    <!--end::Statistics-->
                    <!--begin::Chart-->
                    <div id="kt_card_widget_8_chart" class="min-h-auto" style="height: 125px; min-height: 140px;">
                      <div id="apexchartsc8y3ovgo" class="apexcharts-canvas apexchartsc8y3ovgo apexcharts-theme-light"
                           style="width: 406px; height: 125px;">
                        <svg id="SvgjsSvg1006" width="406" height="125" xmlns="http://www.w3.org/2000/svg" version="1.1"
                             xmlns:xlink="http://www.w3.org/1999/xlink" xmlns:svgjs="http://svgjs.dev"
                             class="apexcharts-svg apexcharts-zoomable hovering-zoom" xmlns:data="ApexChartsNS"
                             transform="translate(0, 0)" style="background: transparent;">
                          <g id="SvgjsG1008" class="apexcharts-inner apexcharts-graphical"
                             transform="translate(-10, 30)">
                            <defs id="SvgjsDefs1007">
                              <clipPath id="gridRectMaskc8y3ovgo">
                                <rect id="SvgjsRect1011" width="441" height="102" x="-3" y="-1" rx="0" ry="0"
                                      opacity="1" stroke-width="0" stroke="none" stroke-dasharray="0"
                                      fill="#fff"></rect>
                              </clipPath>
                              <clipPath id="forecastMaskc8y3ovgo"></clipPath>
                              <clipPath id="nonForecastMaskc8y3ovgo"></clipPath>
                              <clipPath id="gridRectMarkerMaskc8y3ovgo">
                                <rect id="SvgjsRect1012" width="439" height="104" x="-2" y="-2" rx="0" ry="0"
                                      opacity="1" stroke-width="0" stroke="none" stroke-dasharray="0"
                                      fill="#fff"></rect>
                              </clipPath>
                            </defs>
                            <g id="SvgjsG1019" class="apexcharts-xaxis" transform="translate(0, 0)">
                              <g id="SvgjsG1020" class="apexcharts-xaxis-texts-g" transform="translate(0, -4)"></g>
                            </g>
                            <g id="SvgjsG1036" class="apexcharts-grid">
                              <g id="SvgjsG1037" class="apexcharts-gridlines-horizontal">
                                <line id="SvgjsLine1039" x1="0" y1="0" x2="435" y2="0" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                                <line id="SvgjsLine1040" x1="0" y1="25" x2="435" y2="25" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                                <line id="SvgjsLine1041" x1="0" y1="50" x2="435" y2="50" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                                <line id="SvgjsLine1042" x1="0" y1="75" x2="435" y2="75" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                                <line id="SvgjsLine1043" x1="0" y1="100" x2="435" y2="100" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                              </g>
                              <g id="SvgjsG1038" class="apexcharts-gridlines-vertical"></g>
                              <line id="SvgjsLine1045" x1="0" y1="100" x2="435" y2="100" stroke="transparent"
                                    stroke-dasharray="0" stroke-linecap="butt"></line>
                              <line id="SvgjsLine1044" x1="0" y1="1" x2="0" y2="100" stroke="transparent"
                                    stroke-dasharray="0" stroke-linecap="butt"></line>
                            </g>
                            <g id="SvgjsG1013" class="apexcharts-area-series apexcharts-plot-series">
                              <g id="SvgjsG1014" class="apexcharts-series" seriesName="Sales" data:longestSeries="true"
                                 rel="1" data:realIndex="0">
                                <path id="SvgjsPath1017"
                                      d="M0 100L0 37.5C12.6875 37.5 23.5625 7.5 36.25 7.5C48.9375 7.5 59.8125 80 72.5 80C85.1875 80 96.0625 2.5 108.75 2.5C121.4375 2.5 132.3125 45 145 45C157.6875 45 168.5625 10 181.25 10C193.9375 10 204.8125 20 217.5 20C230.1875 20 241.0625 37.5 253.75 37.5C266.4375 37.5 277.3125 2.5 290 2.5C302.6875 2.5 313.5625 37.5 326.25 37.5C338.9375 37.5 349.8125 7.5 362.5 7.5C375.1875 7.5 386.0625 30 398.75 30C411.4375 30 422.3125 7.5 435 7.5C435 7.5 435 7.5 435 100M435 7.5C435 7.5 435 7.5 435 7.5 "
                                      fill="rgba(80,205,137,0)" fill-opacity="1" stroke-opacity="1"
                                      stroke-linecap="butt" stroke-width="0" stroke-dasharray="0"
                                      class="apexcharts-area" index="0" clip-path="url(#gridRectMaskc8y3ovgo)"
                                      pathTo="M 0 100L 0 37.5C 12.6875 37.5 23.5625 7.5 36.25 7.5C 48.9375 7.5 59.8125 80 72.5 80C 85.1875 80 96.0625 2.5 108.75 2.5C 121.4375 2.5 132.3125 45 145 45C 157.6875 45 168.5625 10 181.25 10C 193.9375 10 204.8125 20 217.5 20C 230.1875 20 241.0625 37.5 253.75 37.5C 266.4375 37.5 277.3125 2.5 290 2.5C 302.6875 2.5 313.5625 37.5 326.25 37.5C 338.9375 37.5 349.8125 7.5 362.5 7.5C 375.1875 7.5 386.0625 30 398.75 30C 411.4375 30 422.3125 7.5 435 7.5C 435 7.5 435 7.5 435 100M 435 7.5z"
                                      pathFrom="M -1 150L -1 150L 36.25 150L 72.5 150L 108.75 150L 145 150L 181.25 150L 217.5 150L 253.75 150L 290 150L 326.25 150L 362.5 150L 398.75 150L 435 150"></path>
                                <path id="SvgjsPath1018"
                                      d="M0 37.5C12.6875 37.5 23.5625 7.5 36.25 7.5C48.9375 7.5 59.8125 80 72.5 80C85.1875 80 96.0625 2.5 108.75 2.5C121.4375 2.5 132.3125 45 145 45C157.6875 45 168.5625 10 181.25 10C193.9375 10 204.8125 20 217.5 20C230.1875 20 241.0625 37.5 253.75 37.5C266.4375 37.5 277.3125 2.5 290 2.5C302.6875 2.5 313.5625 37.5 326.25 37.5C338.9375 37.5 349.8125 7.5 362.5 7.5C375.1875 7.5 386.0625 30 398.75 30C411.4375 30 422.3125 7.5 435 7.5C435 7.5 435 7.5 435 7.5 "
                                      fill="none" fill-opacity="1" stroke="#3f4254" stroke-opacity="1"
                                      stroke-linecap="butt" stroke-width="2" stroke-dasharray="0"
                                      class="apexcharts-area" index="0" clip-path="url(#gridRectMaskc8y3ovgo)"
                                      pathTo="M 0 37.5C 12.6875 37.5 23.5625 7.5 36.25 7.5C 48.9375 7.5 59.8125 80 72.5 80C 85.1875 80 96.0625 2.5 108.75 2.5C 121.4375 2.5 132.3125 45 145 45C 157.6875 45 168.5625 10 181.25 10C 193.9375 10 204.8125 20 217.5 20C 230.1875 20 241.0625 37.5 253.75 37.5C 266.4375 37.5 277.3125 2.5 290 2.5C 302.6875 2.5 313.5625 37.5 326.25 37.5C 338.9375 37.5 349.8125 7.5 362.5 7.5C 375.1875 7.5 386.0625 30 398.75 30C 411.4375 30 422.3125 7.5 435 7.5"
                                      pathFrom="M -1 150L -1 150L 36.25 150L 72.5 150L 108.75 150L 145 150L 181.25 150L 217.5 150L 253.75 150L 290 150L 326.25 150L 362.5 150L 398.75 150L 435 150"></path>
                                <g id="SvgjsG1015" class="apexcharts-series-markers-wrap" data:realIndex="0">
                                  <g class="apexcharts-series-markers">
                                    <circle id="SvgjsCircle1053" r="0" cx="253.75" cy="37.5"
                                            class="apexcharts-marker wap30m5lm no-pointer-events" stroke="#3f4254"
                                            fill="#50cd89" fill-opacity="1" stroke-width="2" stroke-opacity="0.9"
                                            default-marker-size="0"></circle>
                                  </g>
                                </g>
                              </g>
                              <g id="SvgjsG1016" class="apexcharts-datalabels" data:realIndex="0"></g>
                            </g>
                            <line id="SvgjsLine1047" x1="253.25" y1="0" x2="253.25" y2="100" stroke="#3f4254"
                                  stroke-dasharray="3" stroke-linecap="butt" class="apexcharts-xcrosshairs" x="253.25"
                                  y="0" width="1" height="100" fill="#b1b9c4" filter="none" fill-opacity="0.9"
                                  stroke-width="1"></line>
                            <line id="SvgjsLine1048" x1="0" y1="0" x2="435" y2="0" stroke="#b6b6b6" stroke-dasharray="0"
                                  stroke-width="1" stroke-linecap="butt" class="apexcharts-ycrosshairs"></line>
                            <line id="SvgjsLine1049" x1="0" y1="0" x2="435" y2="0" stroke-dasharray="0" stroke-width="0"
                                  stroke-linecap="butt" class="apexcharts-ycrosshairs-hidden"></line>
                            <g id="SvgjsG1050" class="apexcharts-yaxis-annotations"></g>
                            <g id="SvgjsG1051" class="apexcharts-xaxis-annotations"></g>
                            <g id="SvgjsG1052" class="apexcharts-point-annotations"></g>
                            <rect id="SvgjsRect1054" width="0" height="0" x="0" y="0" rx="0" ry="0" opacity="1"
                                  stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fefefe"
                                  class="apexcharts-zoom-rect"></rect>
                            <rect id="SvgjsRect1055" width="0" height="0" x="0" y="0" rx="0" ry="0" opacity="1"
                                  stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fefefe"
                                  class="apexcharts-selection-rect"></rect>
                          </g>
                          <g id="SvgjsG1034" class="apexcharts-yaxis" rel="0" transform="translate(-8, 0)">
                            <g id="SvgjsG1035" class="apexcharts-yaxis-texts-g"></g>
                          </g>
                          <rect id="SvgjsRect1046" width="0" height="0" x="0" y="0" rx="0" ry="0" opacity="1"
                                stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fefefe"></rect>
                          <g id="SvgjsG1009" class="apexcharts-annotations"></g>
                        </svg>
                        <div class="apexcharts-legend" style="max-height: 62.5px;"></div>
                        <div class="apexcharts-tooltip apexcharts-theme-light" style="left: 119.891px; top: 40.5px;">
                          <div class="apexcharts-tooltip-title" style="font-family: inherit; font-size: 12px;">Feb 8
                          </div>
                          <div class="apexcharts-tooltip-series-group apexcharts-active"
                               style="order: 1; display: flex;"><span class="apexcharts-tooltip-marker"
                                                                      style="background-color: rgb(80, 205, 137);"></span>
                            <div class="apexcharts-tooltip-text" style="font-family: inherit; font-size: 12px;">
                              <div class="apexcharts-tooltip-y-group"><span class="apexcharts-tooltip-text-y-label">Sales: </span><span
                                  class="apexcharts-tooltip-text-y-value">$4.5K</span></div>
                              <div class="apexcharts-tooltip-goals-group"><span
                                  class="apexcharts-tooltip-text-goals-label"></span><span
                                  class="apexcharts-tooltip-text-goals-value"></span></div>
                              <div class="apexcharts-tooltip-z-group"><span
                                  class="apexcharts-tooltip-text-z-label"></span><span
                                  class="apexcharts-tooltip-text-z-value"></span></div>
                            </div>
                          </div>
                        </div>
                        <div class="apexcharts-xaxistooltip apexcharts-xaxistooltip-bottom apexcharts-theme-light"
                             style="left: 217.359px; top: 132px;">
                          <div class="apexcharts-xaxistooltip-text"
                               style="font-family: inherit; font-size: 12px; min-width: 29.3438px;">Feb 8
                          </div>
                        </div>
                        <div
                            class="apexcharts-yaxistooltip apexcharts-yaxistooltip-0 apexcharts-yaxistooltip-left apexcharts-theme-light">
                          <div class="apexcharts-yaxistooltip-text"></div>
                        </div>
                      </div>
                    </div>
                    <!--end::Chart-->
                  </div>
                  <!--end::Card body-->
                </div>
                <!--end::Card widget 8-->
                <!--begin::Card widget 5-->
                <div class="card card-flush h-md-50 mb-xl-10" style="opacity: 0">
                  <!--begin::Header-->
                  <div class="card-header pt-5">
                    <!--begin::Title-->
                    <div class="card-title d-flex flex-column">
                      <!--begin::Info-->
                      <div class="d-flex align-items-center">
                        <!--begin::Amount-->
                        <span class="fs-2hx fw-bold text-dark me-2 lh-1 ls-n2">1,836</span>
                        <!--end::Amount-->
                        <!--begin::Badge-->
                        <span class="badge badge-light-danger fs-base">
																<!--begin::Svg Icon | path: icons/duotune/arrows/arr065.svg-->
																<span class="svg-icon svg-icon-5 svg-icon-danger ms-n1">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                       xmlns="http://www.w3.org/2000/svg">
																		<rect opacity="0.5" x="11" y="18" width="13" height="2" rx="1"
                                          transform="rotate(-90 11 18)" fill="currentColor"></rect>
																		<path
                                        d="M11.4343 15.4343L7.25 11.25C6.83579 10.8358 6.16421 10.8358 5.75 11.25C5.33579 11.6642 5.33579 12.3358 5.75 12.75L11.2929 18.2929C11.6834 18.6834 12.3166 18.6834 12.7071 18.2929L18.25 12.75C18.6642 12.3358 18.6642 11.6642 18.25 11.25C17.8358 10.8358 17.1642 10.8358 16.75 11.25L12.5657 15.4343C12.2533 15.7467 11.7467 15.7467 11.4343 15.4343Z"
                                        fill="currentColor"></path>
																	</svg>
																</span>
                          <!--end::Svg Icon-->{{dataAdded(this.users)}}</span>
                        <!--end::Badge-->
                      </div>
                      <!--end::Info-->
                      <!--begin::Subtitle-->
                      <span class="text-gray-400 pt-1 fw-semibold fs-6">Total Users</span>
                      <!--end::Subtitle-->
                    </div>
                    <!--end::Title-->
                  </div>
                  <!--end::Header-->
                  <!--begin::Card body-->
                  <div class="card-body d-flex align-items-end pt-0">
                    <!--begin::Progress-->
                    <div class="d-flex align-items-center flex-column mt-3 w-100">
                      <div class="d-flex justify-content-between w-100 mt-auto mb-2">
                        <span class="fw-bolder fs-6 text-dark">1,048 to Goal</span>
                        <span class="fw-bold fs-6 text-gray-400">62%</span>
                      </div>
                      <div class="h-8px mx-3 w-100 bg-light-success rounded">
                        <div class="bg-success rounded h-8px" role="progressbar" style="width: 62%;" aria-valuenow="50"
                             aria-valuemin="0" aria-valuemax="100"></div>
                      </div>
                    </div>
                    <!--end::Progress-->
                  </div>
                  <!--end::Card body-->
                </div>
                <!--end::Card widget 5-->
              </div>
              <!--end::Col-->
              <!--begin::Col-->
              <div class="col-md-6 col-xl-6 mb-xxl-10">
                <!--begin::Card widget 9-->
                <div class="card overflow-hidden h-md-50 mb-5 mb-xl-10">
                  <!--begin::Card body-->
                  <div class="card-body d-flex justify-content-between flex-column px-0 pb-0">
                    <!--begin::Statistics-->
                    <div class="mb-4 px-9">
                      <!--begin::Statistics-->
                      <div class="d-flex align-items-center mb-2">
                        <!--begin::Value-->
                        <span class="fs-2hx fw-bold text-gray-800 me-2 lh-1">{{this.cards.length}}</span>
                        <!--end::Value-->
                        <!--begin::Label-->
                        <span class="badge badge-light-success fs-base">
																<!--begin::Svg Icon | path: icons/duotune/arrows/arr066.svg-->
																<span class="svg-icon svg-icon-5 svg-icon-success ms-n1">
																	<svg width="24" height="24" viewBox="0 0 24 24" fill="none"
                                       xmlns="http://www.w3.org/2000/svg">
																		<rect opacity="0.5" x="13" y="6" width="13" height="2" rx="1"
                                          transform="rotate(90 13 6)" fill="currentColor"></rect>
																		<path
                                        d="M12.5657 8.56569L16.75 12.75C17.1642 13.1642 17.8358 13.1642 18.25 12.75C18.6642 12.3358 18.6642 11.6642 18.25 11.25L12.7071 5.70711C12.3166 5.31658 11.6834 5.31658 11.2929 5.70711L5.75 11.25C5.33579 11.6642 5.33579 12.3358 5.75 12.75C6.16421 13.1642 6.83579 13.1642 7.25 12.75L11.4343 8.56569C11.7467 8.25327 12.2533 8.25327 12.5657 8.56569Z"
                                        fill="currentColor"></path>
																	</svg>
																</span>
                          <!--end::Svg Icon-->{{dataAdded(this.cards)}}</span>
                        <!--end::Label-->
                      </div>
                      <!--end::Statistics-->
                      <!--begin::Description-->
                      <span class="fs-6 fw-semibold text-gray-400">Total Cards</span>
                      <!--end::Description-->
                    </div>
                    <!--end::Statistics-->
                    <!--begin::Chart-->
                    <div id="kt_card_widget_9_chart" class="min-h-auto" style="height: 125px; min-height: 140px;">
                      <div id="apexcharts6vjk2ph4" class="apexcharts-canvas apexcharts6vjk2ph4 apexcharts-theme-light"
                           style="width: 406px; height: 125px;">
                        <svg id="SvgjsSvg1056" width="406" height="125" xmlns="http://www.w3.org/2000/svg" version="1.1"
                             class="apexcharts-svg apexcharts-zoomable hovering-zoom" xmlns:data="ApexChartsNS"
                             transform="translate(0, 0)" style="background: transparent;">
                          <g id="SvgjsG1058" class="apexcharts-inner apexcharts-graphical"
                             transform="translate(-10, 30)">
                            <defs id="SvgjsDefs1057">
                              <clipPath id="gridRectMask6vjk2ph4">
                                <rect id="SvgjsRect1061" width="441" height="102" x="-3" y="-1" rx="0" ry="0"
                                      opacity="1" stroke-width="0" stroke="none" stroke-dasharray="0"
                                      fill="#fff"></rect>
                              </clipPath>
                              <clipPath id="forecastMask6vjk2ph4"></clipPath>
                              <clipPath id="nonForecastMask6vjk2ph4"></clipPath>
                              <clipPath id="gridRectMarkerMask6vjk2ph4">
                                <rect id="SvgjsRect1062" width="439" height="104" x="-2" y="-2" rx="0" ry="0"
                                      opacity="1" stroke-width="0" stroke="none" stroke-dasharray="0"
                                      fill="#fff"></rect>
                              </clipPath>
                            </defs>
                            <g id="SvgjsG1069" class="apexcharts-xaxis" transform="translate(0, 0)">
                              <g id="SvgjsG1070" class="apexcharts-xaxis-texts-g" transform="translate(0, -4)"></g>
                            </g>
                            <g id="SvgjsG1086" class="apexcharts-grid">
                              <g id="SvgjsG1087" class="apexcharts-gridlines-horizontal">
                                <line id="SvgjsLine1089" x1="0" y1="0" x2="435" y2="0" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                                <line id="SvgjsLine1090" x1="0" y1="25" x2="435" y2="25" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                                <line id="SvgjsLine1091" x1="0" y1="50" x2="435" y2="50" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                                <line id="SvgjsLine1092" x1="0" y1="75" x2="435" y2="75" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                                <line id="SvgjsLine1093" x1="0" y1="100" x2="435" y2="100" stroke="#e4e6ef"
                                      stroke-dasharray="4" stroke-linecap="butt" class="apexcharts-gridline"></line>
                              </g>
                              <g id="SvgjsG1088" class="apexcharts-gridlines-vertical"></g>
                              <line id="SvgjsLine1095" x1="0" y1="100" x2="435" y2="100" stroke="transparent"
                                    stroke-dasharray="0" stroke-linecap="butt"></line>
                              <line id="SvgjsLine1094" x1="0" y1="1" x2="0" y2="100" stroke="transparent"
                                    stroke-dasharray="0" stroke-linecap="butt"></line>
                            </g>
                            <g id="SvgjsG1063" class="apexcharts-area-series apexcharts-plot-series">
                              <g id="SvgjsG1064" class="apexcharts-series" seriesName="Visitors"
                                 data:longestSeries="true" rel="1" data:realIndex="0">
                                <path id="SvgjsPath1067"
                                      d="M0 100L0 87.5C12.6875 87.5 23.5625 62.5 36.25 62.5C48.9375 62.5 59.8125 75 72.5 75C85.1875 75 96.0625 50 108.75 50C121.4375 50 132.3125 75 145 75C157.6875 75 168.5625 25 181.25 25C193.9375 25 204.8125 62.5 217.5 62.5C230.1875 62.5 241.0625 75 253.75 75C266.4375 75 277.3125 62.5 290 62.5C302.6875 62.5 313.5625 25 326.25 25C338.9375 25 349.8125 62.5 362.5 62.5C375.1875 62.5 386.0625 12.5 398.75 12.5C411.4375 12.5 422.3125 62.5 435 62.5C435 62.5 435 62.5 435 100M435 62.5C435 62.5 435 62.5 435 62.5 "
                                      fill="rgba(80,205,137,0)" fill-opacity="1" stroke-opacity="1"
                                      stroke-linecap="butt" stroke-width="0" stroke-dasharray="0"
                                      class="apexcharts-area" index="0" clip-path="url(#gridRectMask6vjk2ph4)"
                                      pathTo="M 0 100L 0 87.5C 12.6875 87.5 23.5625 62.5 36.25 62.5C 48.9375 62.5 59.8125 75 72.5 75C 85.1875 75 96.0625 50 108.75 50C 121.4375 50 132.3125 75 145 75C 157.6875 75 168.5625 25 181.25 25C 193.9375 25 204.8125 62.5 217.5 62.5C 230.1875 62.5 241.0625 75 253.75 75C 266.4375 75 277.3125 62.5 290 62.5C 302.6875 62.5 313.5625 25 326.25 25C 338.9375 25 349.8125 62.5 362.5 62.5C 375.1875 62.5 386.0625 12.5 398.75 12.5C 411.4375 12.5 422.3125 62.5 435 62.5C 435 62.5 435 62.5 435 100M 435 62.5z"
                                      pathFrom="M -1 125L -1 125L 36.25 125L 72.5 125L 108.75 125L 145 125L 181.25 125L 217.5 125L 253.75 125L 290 125L 326.25 125L 362.5 125L 398.75 125L 435 125"></path>
                                <path id="SvgjsPath1068"
                                      d="M0 87.5C12.6875 87.5 23.5625 62.5 36.25 62.5C48.9375 62.5 59.8125 75 72.5 75C85.1875 75 96.0625 50 108.75 50C121.4375 50 132.3125 75 145 75C157.6875 75 168.5625 25 181.25 25C193.9375 25 204.8125 62.5 217.5 62.5C230.1875 62.5 241.0625 75 253.75 75C266.4375 75 277.3125 62.5 290 62.5C302.6875 62.5 313.5625 25 326.25 25C338.9375 25 349.8125 62.5 362.5 62.5C375.1875 62.5 386.0625 12.5 398.75 12.5C411.4375 12.5 422.3125 62.5 435 62.5C435 62.5 435 62.5 435 62.5 "
                                      fill="none" fill-opacity="1" stroke="#3f4254" stroke-opacity="1"
                                      stroke-linecap="butt" stroke-width="2" stroke-dasharray="0"
                                      class="apexcharts-area" index="0" clip-path="url(#gridRectMask6vjk2ph4)"
                                      pathTo="M 0 87.5C 12.6875 87.5 23.5625 62.5 36.25 62.5C 48.9375 62.5 59.8125 75 72.5 75C 85.1875 75 96.0625 50 108.75 50C 121.4375 50 132.3125 75 145 75C 157.6875 75 168.5625 25 181.25 25C 193.9375 25 204.8125 62.5 217.5 62.5C 230.1875 62.5 241.0625 75 253.75 75C 266.4375 75 277.3125 62.5 290 62.5C 302.6875 62.5 313.5625 25 326.25 25C 338.9375 25 349.8125 62.5 362.5 62.5C 375.1875 62.5 386.0625 12.5 398.75 12.5C 411.4375 12.5 422.3125 62.5 435 62.5"
                                      pathFrom="M -1 125L -1 125L 36.25 125L 72.5 125L 108.75 125L 145 125L 181.25 125L 217.5 125L 253.75 125L 290 125L 326.25 125L 362.5 125L 398.75 125L 435 125"></path>
                                <g id="SvgjsG1065" class="apexcharts-series-markers-wrap" data:realIndex="0">
                                  <g class="apexcharts-series-markers">
                                    <circle id="SvgjsCircle1103" r="0" cx="36.25" cy="62.5"
                                            class="apexcharts-marker w5bcqrqvy no-pointer-events" stroke="#3f4254"
                                            fill="#50cd89" fill-opacity="1" stroke-width="2" stroke-opacity="0.9"
                                            default-marker-size="0"></circle>
                                  </g>
                                </g>
                              </g>
                              <g id="SvgjsG1066" class="apexcharts-datalabels" data:realIndex="0"></g>
                            </g>
                            <line id="SvgjsLine1097" x1="35.75" y1="0" x2="35.75" y2="100" stroke="#3f4254"
                                  stroke-dasharray="3" stroke-linecap="butt" class="apexcharts-xcrosshairs" x="35.75"
                                  y="0" width="1" height="100" fill="#b1b9c4" filter="none" fill-opacity="0.9"
                                  stroke-width="1"></line>
                            <line id="SvgjsLine1098" x1="0" y1="0" x2="435" y2="0" stroke="#b6b6b6" stroke-dasharray="0"
                                  stroke-width="1" stroke-linecap="butt" class="apexcharts-ycrosshairs"></line>
                            <line id="SvgjsLine1099" x1="0" y1="0" x2="435" y2="0" stroke-dasharray="0" stroke-width="0"
                                  stroke-linecap="butt" class="apexcharts-ycrosshairs-hidden"></line>
                            <g id="SvgjsG1100" class="apexcharts-yaxis-annotations"></g>
                            <g id="SvgjsG1101" class="apexcharts-xaxis-annotations"></g>
                            <g id="SvgjsG1102" class="apexcharts-point-annotations"></g>
                            <rect id="SvgjsRect1104" width="0" height="0" x="0" y="0" rx="0" ry="0" opacity="1"
                                  stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fefefe"
                                  class="apexcharts-zoom-rect"></rect>
                            <rect id="SvgjsRect1105" width="0" height="0" x="0" y="0" rx="0" ry="0" opacity="1"
                                  stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fefefe"
                                  class="apexcharts-selection-rect"></rect>
                          </g>
                          <g id="SvgjsG1084" class="apexcharts-yaxis" rel="0" transform="translate(-8, 0)">
                            <g id="SvgjsG1085" class="apexcharts-yaxis-texts-g"></g>
                          </g>
                          <rect id="SvgjsRect1096" width="0" height="0" x="0" y="0" rx="0" ry="0" opacity="1"
                                stroke-width="0" stroke="none" stroke-dasharray="0" fill="#fefefe"></rect>
                          <g id="SvgjsG1059" class="apexcharts-annotations"></g>
                        </svg>
                        <div class="apexcharts-legend" style="max-height: 62.5px;"></div>
                        <div class="apexcharts-tooltip apexcharts-theme-light" style="left: 37.25px; top: 65.5px;">
                          <div class="apexcharts-tooltip-title" style="font-family: inherit; font-size: 12px;">Feb 2
                          </div>
                          <div class="apexcharts-tooltip-series-group apexcharts-active"
                               style="order: 1; display: flex;"><span class="apexcharts-tooltip-marker"
                                                                      style="background-color: rgb(80, 205, 137);"></span>
                            <div class="apexcharts-tooltip-text" style="font-family: inherit; font-size: 12px;">
                              <div class="apexcharts-tooltip-y-group"><span class="apexcharts-tooltip-text-y-label">Visitors: </span><span
                                  class="apexcharts-tooltip-text-y-value">2.5K</span></div>
                              <div class="apexcharts-tooltip-goals-group"><span
                                  class="apexcharts-tooltip-text-goals-label"></span><span
                                  class="apexcharts-tooltip-text-goals-value"></span></div>
                              <div class="apexcharts-tooltip-z-group"><span
                                  class="apexcharts-tooltip-text-z-label"></span><span
                                  class="apexcharts-tooltip-text-z-value"></span></div>
                            </div>
                          </div>
                        </div>
                        <div class="apexcharts-xaxistooltip apexcharts-xaxistooltip-bottom apexcharts-theme-light"
                             style="left: -0.0703125px; top: 132px;">
                          <div class="apexcharts-xaxistooltip-text"
                               style="font-family: inherit; font-size: 12px; min-width: 29.3438px;">Feb 2
                          </div>
                        </div>
                        <div
                            class="apexcharts-yaxistooltip apexcharts-yaxistooltip-0 apexcharts-yaxistooltip-left apexcharts-theme-light">
                          <div class="apexcharts-yaxistooltip-text"></div>
                        </div>
                      </div>
                    </div>
                    <!--end::Chart-->
                  </div>
                  <!--end::Card body-->
                </div>
                <!--end::Card widget 9-->

              </div>
              <!--end::Col-->
            </div>
            <!--end::Row-->
          </div>
          <!--end::Col-->

        </div>
        <!--end::Row-->
      </div>
      <!--end::Post-->
      <!--      &lt;!&ndash;begin::Footer&ndash;&gt;-->
      <!--      <div class="footer py-4 d-flex flex-column flex-md-row flex-stack" id="kt_footer">-->
      <!--        &lt;!&ndash;begin::Copyright&ndash;&gt;-->
      <!--        <div class="text-dark order-2 order-md-1">-->
      <!--          <span class="text-muted fw-semibold me-1">2022©</span>-->
      <!--          <a href="" target="_blank" class="text-gray-800 text-hover-primary">Keenthemes</a>-->
      <!--        </div>-->
      <!--        &lt;!&ndash;end::Copyright&ndash;&gt;-->
      <!--        &lt;!&ndash;begin::Menu&ndash;&gt;-->
      <!--        <ul class="menu menu-gray-600 menu-hover-primary fw-semibold order-1">-->
      <!--          <li class="menu-item">-->
      <!--            <a href="" target="_blank" class="menu-link px-2">About</a>-->
      <!--          </li>-->
      <!--          <li class="menu-item">-->
      <!--            <a href="" target="_blank" class="menu-link px-2">Support</a>-->
      <!--          </li>-->
      <!--          <li class="menu-item">-->
      <!--            <a href="" target="_blank" class="menu-link px-2">Purchase</a>-->
      <!--          </li>-->
      <!--        </ul>-->
      <!--        &lt;!&ndash;end::Menu&ndash;&gt;-->
      <!--      </div>-->
      <!--      &lt;!&ndash;end::Footer&ndash;&gt;-->
    </div>
  </div>
</template>

<script>
import {defineComponent} from 'vue'
import {mapGetters} from "vuex";

export default defineComponent({
  name: "mainPage",
  methods: {
    getData() {
      if (this.loggedIn && this.users.length === 0) {
        this.$store.dispatch("admin/getUsersMethod")
      }
      if (this.loggedIn && !this.cardsLoaded) {
        this.$store.dispatch("admin/getCardsMethod");
      }
      return true
    },
    dataAdded(obj) {
      let newData = []
      let currentDate = new Date().toLocaleDateString()
      let dataArr = Object.values(obj);
      for (const item of dataArr) {
        let currentDateReverse = currentDate.split(".")[2] + "-" + currentDate.split(".")[1] + "-" + currentDate.split(".")[0]
        if(item.CreatedAt.split("T")[0] === currentDateReverse) {
          newData.push(item)
        }
      }
      return newData.length
    }
  },
  computed: {
    ...mapGetters({
      loggedIn: "auth/isLoggedIn",
      users: "admin/getUsers",
      cards: "admin/getCards",
      cardsLoaded: "admin/getCardsLoaded",
      user: "auth/currentUser"
    }),
    currentUser() {
      const user = this.$store.getters["auth/currentUser"];
      if (user) {
        return user.email;
      }
      return "";
    },
  },
})
</script>