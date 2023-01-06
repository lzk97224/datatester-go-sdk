/**
 * Apache 2.0
 * Copyright 2022 Beijing Volcano Engine Technology Co., Ltd.
 */

package mock

const MetaMock = "{\"experiments\":{\"77764\":{\"id\":\"77764\",\"name\":\"2656e0f545101eadc5bbb905b562e522\",\"raw_name\":\"过滤参数测试实验01\",\"layer_id\":\"60196\",\"status\":1,\"variants\":{\"120295\":{\"id\":\"120295\",\"entity_id\":\"77764\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"filter_param\":{\"id\":\"\",\"value\":\"a\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120296\":{\"id\":\"120296\",\"entity_id\":\"77764\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"filter_param\":{\"id\":\"\",\"value\":\"b\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"bool_param\",\"op\":\"in\",\"logic_operator\":\"&&\",\"value\":[true],\"type\":\"boolean\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"bool_param\",\"op\":\"in\",\"logic_operator\":\"&&\",\"value\":[false],\"type\":\"boolean\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"number_param\",\"op\":\"in\",\"logic_operator\":\"&&\",\"value\":[12345],\"type\":\"number\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"number_param\",\"op\":\">\",\"logic_operator\":\"&&\",\"value\":67890,\"type\":\"number\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"number_param\",\"op\":\"<\",\"logic_operator\":\"&&\",\"value\":5678,\"type\":\"number\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\"in\",\"logic_operator\":\"&&\",\"value\":[\"str1\"],\"type\":\"string\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\"is_null\",\"logic_operator\":\"&&\",\"value\":\"is_null\",\"type\":\"string\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\">\",\"logic_operator\":\"&&\",\"value\":\"jkl\",\"type\":\"string\",\"method\":\"dict\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\"<\",\"logic_operator\":\"&&\",\"value\":\"efg\",\"type\":\"string\",\"method\":\"dict\",\"property_type\":\"\",\"editor_url\":\"\"}]}],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120295\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120296\"}]},\"white_list\":{},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669012339,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":0,\"father_experiment_id\":\"\",\"modify_time\":1669013233,\"associated_relations\":[]},\"77767\":{\"id\":\"77767\",\"name\":\"e4e1e87f90e52ead2621d2360540849e\",\"raw_name\":\"过滤参数测试实验02\",\"layer_id\":\"60197\",\"status\":1,\"variants\":{\"120297\":{\"id\":\"120297\",\"entity_id\":\"77767\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"filter_param_02\":{\"id\":\"\",\"value\":\"a\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120298\":{\"id\":\"120298\",\"entity_id\":\"77767\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"filter_param_02\":{\"id\":\"\",\"value\":\"b\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"bool_param\",\"op\":\"ni\",\"logic_operator\":\"&&\",\"value\":[true],\"type\":\"boolean\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"bool_param\",\"op\":\"ni\",\"logic_operator\":\"&&\",\"value\":[false],\"type\":\"boolean\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"number_param\",\"op\":\"ni\",\"logic_operator\":\"&&\",\"value\":[12345],\"type\":\"number\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\">=\",\"logic_operator\":\"&&\",\"value\":\"5.6.7\",\"type\":\"string\",\"method\":\"version\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\"<=\",\"logic_operator\":\"&&\",\"value\":\"2.3.4\",\"type\":\"string\",\"method\":\"version\",\"property_type\":\"\",\"editor_url\":\"\"}]}],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120297\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120298\"}]},\"white_list\":{},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669012581,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":0,\"father_experiment_id\":\"\",\"modify_time\":1669013591,\"associated_relations\":[]},\"77768\":{\"id\":\"77768\",\"name\":\"cb7334f83db9d0025ff2672b09927bd8\",\"raw_name\":\"过滤参数测试实验03\",\"layer_id\":\"60198\",\"status\":1,\"variants\":{\"120299\":{\"id\":\"120299\",\"entity_id\":\"77768\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"filter_param_03\":{\"id\":\"\",\"value\":\"a\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120300\":{\"id\":\"120300\",\"entity_id\":\"77768\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"filter_param_03\":{\"id\":\"\",\"value\":\"b\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"bool_param\",\"op\":\"in\",\"logic_operator\":\"&&\",\"value\":[true,false],\"type\":\"boolean\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"number_param\",\"op\":\">=\",\"logic_operator\":\"&&\",\"value\":345.56,\"type\":\"number\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"number_param\",\"op\":\"<=\",\"logic_operator\":\"&&\",\"value\":123.45,\"type\":\"number\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]},{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\"ni\",\"logic_operator\":\"&&\",\"value\":[\"str1\",\"str2\"],\"type\":\"string\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]}],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120299\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120300\"}]},\"white_list\":{},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669012857,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":0,\"father_experiment_id\":\"\",\"modify_time\":1669013598,\"associated_relations\":[]},\"77769\":{\"id\":\"77769\",\"name\":\"df60d17c66da9dccc8a1ee3c1befcfb4\",\"raw_name\":\"过滤参数测试实验04\",\"layer_id\":\"60200\",\"status\":1,\"variants\":{\"120301\":{\"id\":\"120301\",\"entity_id\":\"77769\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"filter_param_04\":{\"id\":\"\",\"value\":\"a\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120302\":{\"id\":\"120302\",\"entity_id\":\"77769\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"filter_param_04\":{\"id\":\"\",\"value\":\"b\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\"is_not_null\",\"logic_operator\":\"&&\",\"value\":\"is_not_null\",\"type\":\"string\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]}],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120301\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120302\"}]},\"white_list\":{},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669013174,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":0,\"father_experiment_id\":\"\",\"modify_time\":1669013577,\"associated_relations\":[]},\"77770\":{\"id\":\"77770\",\"name\":\"3dc26716d8a230f335f187301781d6b8\",\"raw_name\":\"白名单实验-不命中过滤条件\",\"layer_id\":\"60202\",\"status\":3,\"variants\":{\"120303\":{\"id\":\"120303\",\"entity_id\":\"77770\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"allow_without_filter\":{\"id\":\"\",\"value\":true,\"type\":\"bool\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120304\":{\"id\":\"120304\",\"entity_id\":\"77770\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"allow_without_filter\":{\"id\":\"\",\"value\":false,\"type\":\"bool\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\"is_not_null\",\"logic_operator\":\"&&\",\"value\":\"is_not_null\",\"type\":\"string\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]}],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120303\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120304\"}]},\"white_list\":{\"test_user\":\"120303\",\"test_user_02\":\"120304\"},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669014030,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":0,\"father_experiment_id\":\"\",\"modify_time\":1669014160,\"associated_relations\":[]},\"77771\":{\"id\":\"77771\",\"name\":\"bbb245a70d3a4edcab694a42cc1d5cd0\",\"raw_name\":\"白名单实验-命中过滤条件\",\"layer_id\":\"60203\",\"status\":3,\"variants\":{\"120305\":{\"id\":\"120305\",\"entity_id\":\"77771\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"allow_with_filter\":{\"id\":\"\",\"value\":0,\"type\":\"double\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120306\":{\"id\":\"120306\",\"entity_id\":\"77771\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"allow_with_filter\":{\"id\":\"\",\"value\":1,\"type\":\"double\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"str_param\",\"op\":\"is_not_null\",\"logic_operator\":\"&&\",\"value\":\"is_not_null\",\"type\":\"string\",\"method\":\"\",\"property_type\":\"\",\"editor_url\":\"\"}]}],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120305\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120306\"}]},\"white_list\":{\"test_user\":\"120305\",\"test_user_02\":\"120306\"},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669014116,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":1,\"father_experiment_id\":\"\",\"modify_time\":1669014118,\"associated_relations\":[]},\"77772\":{\"id\":\"77772\",\"name\":\"f2e878951cf1b72f8c22b613199bdcd6\",\"raw_name\":\"冻结实验\",\"layer_id\":\"60204\",\"status\":1,\"variants\":{\"120307\":{\"id\":\"120307\",\"entity_id\":\"77772\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"freeze\":{\"id\":\"\",\"value\":{\"freeze\":0},\"type\":\"json\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120308\":{\"id\":\"120308\",\"entity_id\":\"77772\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"freeze\":{\"id\":\"\",\"value\":{\"freeze\":1},\"type\":\"json\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120307\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120308\"}]},\"white_list\":{\"test_user\":\"120307\",\"test_user_02\":\"120308\"},\"immediately\":1,\"freeze_status\":1,\"launch_start_time\":1669014288,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":1,\"father_experiment_id\":\"\",\"modify_time\":1669014321,\"associated_relations\":[]},\"77773\":{\"id\":\"77773\",\"name\":\"4230d9f68bb65ecb5b375b12538f6e53\",\"raw_name\":\"父实验01\",\"layer_id\":\"60205\",\"status\":1,\"variants\":{\"120309\":{\"id\":\"120309\",\"entity_id\":\"77773\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"father\":{\"id\":\"\",\"value\":\"a\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120310\":{\"id\":\"120310\",\"entity_id\":\"77773\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"father\":{\"id\":\"\",\"value\":\"b\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120309\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120310\"}]},\"white_list\":{},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669014576,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":1,\"father_experiment_id\":\"\",\"modify_time\":1669014580,\"associated_relations\":[]},\"77774\":{\"id\":\"77774\",\"name\":\"40f1b755b7e5e8ccced854b782b887b3\",\"raw_name\":\"子实验01\",\"layer_id\":\"60206\",\"status\":1,\"variants\":{\"120311\":{\"id\":\"120311\",\"entity_id\":\"77774\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"child_01\":{\"id\":\"\",\"value\":\"a\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":[\"120309\"],\"disabled\":0},\"120312\":{\"id\":\"120312\",\"entity_id\":\"77774\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"child_01\":{\"id\":\"\",\"value\":\"b\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":[\"120310\"],\"disabled\":0},\"120313\":{\"id\":\"120313\",\"entity_id\":\"77774\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本2\",\"config\":{\"child_01\":{\"id\":\"\",\"value\":\"c\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[],\"traffic_allocation\":[{\"begin\":0,\"end\":333,\"entity_id\":\"120311\"},{\"begin\":333,\"end\":666,\"entity_id\":\"120312\"},{\"begin\":666,\"end\":1000,\"entity_id\":\"120313\"}]},\"white_list\":{},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669014711,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":0,\"father_experiment_id\":\"77773\",\"modify_time\":1669014716,\"associated_relations\":[]},\"77775\":{\"id\":\"77775\",\"name\":\"842b2451d1fde3c99ab594a1ceae9c56\",\"raw_name\":\"关联父实验\",\"layer_id\":\"60207\",\"status\":1,\"variants\":{\"120314\":{\"id\":\"120314\",\"entity_id\":\"77775\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"asso\":{\"id\":\"\",\"value\":\"a\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120315\":{\"id\":\"120315\",\"entity_id\":\"77775\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"asso\":{\"id\":\"\",\"value\":\"b\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120314\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120315\"}]},\"white_list\":{},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669014842,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":0,\"father_experiment_id\":\"\",\"modify_time\":1669014846,\"associated_relations\":[]},\"77779\":{\"id\":\"77779\",\"name\":\"81b77a493a890341af004e1f90273d3a\",\"raw_name\":\"关联子实验\",\"layer_id\":\"60213\",\"status\":1,\"variants\":{\"120323\":{\"id\":\"120323\",\"entity_id\":\"77779\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"asso_child\":{\"id\":\"\",\"value\":\"a\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120324\":{\"id\":\"120324\",\"entity_id\":\"77779\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"asso_child\":{\"id\":\"\",\"value\":\"b\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"experiment_cohort\",\"op\":\"in\",\"logic_operator\":\"&&\",\"value\":[77775],\"type\":\"int\",\"method\":\"\",\"property_type\":\"experiment_cohort\",\"editor_url\":\"\"}]}],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120323\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120324\"}]},\"white_list\":{},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669017283,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":0,\"father_experiment_id\":\"\",\"modify_time\":1669017287,\"associated_relations\":[\"77775\"]},\"77780\":{\"id\":\"77780\",\"name\":\"177575c5d4af39b8cc2eb1f9f75ef750\",\"raw_name\":\"关联子实验_02\",\"layer_id\":\"60214\",\"status\":1,\"variants\":{\"120325\":{\"id\":\"120325\",\"entity_id\":\"77780\",\"type\":0,\"match_mode\":\"\",\"name\":\"对照版本\",\"config\":{\"asso_child_02\":{\"id\":\"\",\"value\":\"a\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0},\"120326\":{\"id\":\"120326\",\"entity_id\":\"77780\",\"type\":1,\"match_mode\":\"\",\"name\":\"实验版本1\",\"config\":{\"asso_child_02\":{\"id\":\"\",\"value\":\"b\",\"type\":\"string\"}},\"config_for_pages\":{},\"father_variants\":null,\"disabled\":0}},\"release\":{\"filter\":[{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"experiment_cohort\",\"op\":\"ni\",\"logic_operator\":\"&&\",\"value\":[77775],\"type\":\"int\",\"method\":\"\",\"property_type\":\"experiment_cohort\",\"editor_url\":\"\"}]}],\"traffic_allocation\":[{\"begin\":0,\"end\":500,\"entity_id\":\"120325\"},{\"begin\":500,\"end\":1000,\"entity_id\":\"120326\"}]},\"white_list\":{},\"immediately\":1,\"freeze_status\":0,\"launch_start_time\":1669018277,\"version_freeze_status\":1,\"save_param\":1,\"experiment_mode\":1,\"filter_allowlist\":0,\"father_experiment_id\":\"\",\"modify_time\":1669018290,\"associated_relations\":[\"77775\"]}},\"layers\":{\"60196\":{\"id\":\"60196\",\"name\":\"ee3b745ba765b64b67f750f03179a63ee2695b77\",\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"77764\"}],\"experiment_ids\":[\"77764\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669012278},\"60197\":{\"id\":\"60197\",\"name\":\"0fbcfa53953933c879538eb1f93b1d8cbad8baaa\",\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"77767\"}],\"experiment_ids\":[\"77767\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669012559},\"60198\":{\"id\":\"60198\",\"name\":\"beecb7a5c3e4018cbbae70aee3421d07064cbe5e\",\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"77768\"}],\"experiment_ids\":[\"77768\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669012839},\"60200\":{\"id\":\"60200\",\"name\":\"3b42ef42675c29d03c2442e523cbe11fcc055585\",\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"77769\"}],\"experiment_ids\":[\"77769\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669013155},\"60202\":{\"id\":\"60202\",\"name\":\"ed71ee668884a0cccb00dda338cd66cbc27ef880\",\"traffic_allocation\":[],\"experiment_ids\":[\"77770\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669014034},\"60203\":{\"id\":\"60203\",\"name\":\"98467e8699845c3025fc3881b92e9887368f4f8f\",\"traffic_allocation\":[],\"experiment_ids\":[\"77771\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669014118},\"60204\":{\"id\":\"60204\",\"name\":\"13cad586b1a29df4c9646973c622a9c7bc755c7b\",\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"77772\"}],\"experiment_ids\":[\"77772\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669014271},\"60205\":{\"id\":\"60205\",\"name\":\"5be01f1be7967df0e44d593c96a8988030f9dec0\",\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"77773\"}],\"experiment_ids\":[\"77773\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669014416},\"60206\":{\"id\":\"60206\",\"name\":\"4d822ed38ca362205072023239d2d0c457a77e10\",\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"77774\"}],\"experiment_ids\":[\"77774\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669014694},\"60207\":{\"id\":\"60207\",\"name\":\"1233c60769d809a84e44912615dde72a7e2757ce\",\"traffic_allocation\":[{\"begin\":0,\"end\":400,\"entity_id\":\"77775\"}],\"experiment_ids\":[\"77775\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669014821},\"60213\":{\"id\":\"60213\",\"name\":\"121f10bf810caece88cd5a6d772d611e87eac96d\",\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"77779\"}],\"experiment_ids\":[\"77779\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669017259},\"60214\":{\"id\":\"60214\",\"name\":\"beb3c2f23f0860489f6362570953dc82092d7888\",\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"77780\"}],\"experiment_ids\":[\"77780\"],\"hash_strategy\":\"ssid\",\"modify_time\":1669017489}},\"features\":{\"10202\":{\"id\":\"10202\",\"name\":\"服务端feature\",\"side_type\":\"server\",\"releases\":[{\"filter\":[{\"id\":\"\",\"logic_operator\":\"||\",\"conditions\":[{\"key\":\"age\",\"op\":\">=\",\"logic_operator\":\"&&\",\"value\":12345,\"type\":\"number\",\"method\":\"\",\"property_type\":\"common_param\",\"editor_url\":\"\"}]}],\"traffic_allocation\":[{\"begin\":0,\"end\":1000,\"entity_id\":\"20101622\"},{\"begin\":1000,\"end\":1000,\"entity_id\":\"20101623\"}]},{\"filter\":[],\"traffic_allocation\":[{\"begin\":0,\"end\":0,\"entity_id\":\"20101622\"},{\"begin\":0,\"end\":1000,\"entity_id\":\"20101623\"}]}],\"status\":1,\"launch_start_time\":1669017782,\"white_list\":{\"test_user\":\"20101622\",\"test_user_02\":\"20101623\"},\"variants\":{\"20101622\":{\"id\":\"20101622\",\"entity_id\":\"10202\",\"type\":0,\"match_mode\":\"\",\"name\":\"变体1\",\"config\":{\"feature\":{\"id\":\"\",\"value\":true,\"type\":\"bool\"}},\"config_for_pages\":null,\"father_variants\":null,\"disabled\":0},\"20101623\":{\"id\":\"20101623\",\"entity_id\":\"10202\",\"type\":0,\"match_mode\":\"\",\"name\":\"变体2\",\"config\":{\"feature\":{\"id\":\"\",\"value\":false,\"type\":\"bool\"}},\"config_for_pages\":null,\"father_variants\":null,\"disabled\":0}},\"modify_time\":1669017783}},\"hash_strategy\":\"ssid\",\"modify_time\":1669018290,\"app_id\":\"176894\",\"filter_keys\":null,\"total_white_list\":null}"
