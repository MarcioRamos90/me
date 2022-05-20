#!/bin/python3

import requests


def clean_input_arr(message):
	arr = input(message).split(',') 
	if len(arr) == 1 and arr[0] == "":
		arr = []
	return arr


url = input("Enter the url to retrieve: ")
cookies_arr = clean_input_arr("set cookies: ")
headers_arr = clean_input_arr("set headers: ")

cookies = {}

print('-->', cookies_arr, headers_arr)
for item in cookies_arr:
    key, value = item.split("=")
    cookies[key] = value

print('C -->', cookies)

headers = {}

for item in headers_arr:
    key, value = item.split("=")
    headers[key] = value

print('H -->', headers)

response = requests.get(url, cookies=cookies, headers=headers)

print(response.request.headers)
print(response.text)

