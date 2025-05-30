import requests
import time
import sys

if len(sys.argv) > 1:
    no_of_items = int(sys.argv[1])
else:
    no_of_items = 100

def print_times(header, total_time):
    print()
    print(f'{header} average time per request:')
    print(f'{round(total_time / no_of_items * 1000, 2)} ms')

# populate data
start = time.time()

print(f"Testing with {no_of_items} items")

for i in range(no_of_items):
    res = requests.put(f'http://localhost:8080/{i}', json={"value": i})
    if res.status_code != 200:
        print("put request failed")
        exit(1)

end = time.time()

put_time_total = end - start

print_times("PUT", put_time_total)

# get data
start = time.time()

for i in range(no_of_items):
    res = requests.get(f'http://localhost:8080/{i}')
    if res.status_code != 200:
        print("get request failed")
        exit(1)

end = time.time()

get_time_total = end - start

print_times("GET", get_time_total)

# delete data
for i in range(no_of_items):
    res = requests.delete(f'http://localhost:8080/{i}')
    if res.status_code != 200:
        print("delete request failed")
        exit(1)

print("items deleted")

