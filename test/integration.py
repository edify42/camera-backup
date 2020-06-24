import sys, hashlib

def calculate_s3_etag(file_path, chunk_size=8 * 1024 * 1024):
    md5s = []

    with open(file_path, 'rb') as fp:
        while True:
            data = fp.read(chunk_size)
            if not data:
                break
            md5s.append(hashlib.md5(data))

    if len(md5s) < 1:
        return '{}'.format(hashlib.md5().hexdigest())

    if len(md5s) == 1:
        return '{}'.format(md5s[0].hexdigest())

    digests = b''.join(m.digest() for m in md5s)
    digests_md5 = hashlib.md5(digests)
    return '{}-{}'.format(digests_md5.hexdigest(), len(md5s))

if __name__ == '__main__':
	if  len(sys.argv) < 2 or len(sys.argv) > 3:

		print("python %s source_path" % (sys.argv[0]))
		exit()

	source_path = sys.argv[1]
	chunk_size = 8 # sys.argv[2] # Yeah we don't need to be clever here.
	chunk_size = int(chunk_size) * 1024 * 1024
	print(calculate_s3_etag(source_path,chunk_size))
    