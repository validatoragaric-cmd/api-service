# utils.py

import logging
from datetime import datetime

def get_current_timestamp():
    return datetime.now().strftime("%Y-%m-%d %H:%M:%S")

def get_random_string(length):
    import secrets
    import string
    return ''.join(secrets.choice(string.ascii_letters + string.digits) for _ in range(length))

def log_error(message, error):
    logging.error(f"[{get_current_timestamp()}] {message} - {error}")

def log_info(message):
    logging.info(f"[{get_current_timestamp()}] {message}")