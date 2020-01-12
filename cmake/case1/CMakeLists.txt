cmake_minimum_required(VERSION 2.6)

project(case1)

include_directories(${PROJECT_SOURCE_DIR})

add_subdirectory(sub1)
add_subdirectory(sub2)

add_subdirectory(main1)
