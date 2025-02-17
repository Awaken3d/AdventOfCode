import hashlib
import threading
import queue

q = queue.Queue(30)
q_hash = queue.Queue(30)
results = []
input : str = "iwrupvqb"

def calculate_hashes():
    try:
        while not q.is_shutdown:
            number_range = q.get(True, None)
            for i in range(number_range - 100, number_range + 1):
                input_cypher = input + str(i)
                hex_result = hashlib.md5(input_cypher.encode()).hexdigest()
                q_hash.put((hex_result, input_cypher), True, None)
    except queue.ShutDown:
        print("calculators ending")

def parse_hashes():
    try:
        while not q_hash.is_shutdown:
            hash_tuple = q_hash.get(True, None)
            if all(char == '0' for char in hash_tuple[0][:5]) and not results:
                results.append(hash_tuple[1])
                print("result is " + hash_tuple[1])
    except queue.ShutDown:
        print("parsers ending")

def main():
    for i in range(5):
        threading.Thread(target = calculate_hashes).start()
        threading.Thread(target = parse_hashes).start()

    for i in range(100, 10000000000, 100):
        if results:
            break
        print("adding new batch " + str(i))
        q.put(i, True, None)
    
    print("result found")
    print(results)
    q.shutdown()
    q_hash.shutdown()

if __name__ == "__main__":
    main()
