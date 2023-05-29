import socket
import ssl


# ===================> Server <===================

def start_tls_server(key: str, cert: str, client_crt: str, port: int, host: str):
    """"
    Starts a TLS server on the given port and host.
    
    Args:
        key (str): The server's private key file.
        cert (str): The server's certificate file.
        client_crt (str): The client's certificate file.
        port (int): The server's port.
        host (str): The server's host.
        
    Returns:
        None
    """
    # Create a socket and bind it to the host and port
    server_socket = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    server_socket.bind((host, port))

    # Create an SSL context
    context = ssl.create_default_context(ssl.Purpose.CLIENT_AUTH)
    context.load_cert_chain(certfile=cert, keyfile=key)
    context.load_verify_locations(cafile=client_crt)

    # Listen for incoming connections forever
    server_socket.listen()
    print(f'[+] Listening as {host}:{port}')
    while True:
        print('Waiting for a client connection...')
        # Accept a client socket
        client_socket, client_address = server_socket.accept()
        print(f'[+] {client_address} connected.')
        # Handle the client connection in a separate thread
        handle_client(client_socket, context)

def handle_client(client_socket: socket.socket, context: ssl.SSLContext):
    """Handle a client connection.

    Args:
        client_socket (socket.socket): The client socket.

    Returns:
        None
    """
    # Perform SSL handshake
    ssl_socket = context.wrap_socket(client_socket, server_side=True)
    # Receive data from the client
    encrypted_data = ssl_socket.recv(1024)
    data = encrypted_data.decode('utf-8')
    print(f'[+] Received: {data}')
    # Send data to the client
    response = b"Hello from the server!"
    ssl_socket.send(response)
    print(f'[+] Sent: {response}')
    # Close the socket
    ssl_socket.close()


if __name__ == '__main__':
    # Server's host and port
    HOST = 'localhost'
    PORT = 12345
    # Server's certificate and private key files
    CERT_FILE = '../certs/server.crt'
    KEY_FILE = '../certs/server.key'
    CLIENT_CERT = '../certs/client.crt'
    start_tls_server(KEY_FILE, CERT_FILE, CLIENT_CERT, PORT, HOST)