import logging

logger = logging.getLogger("debug")
logger.setLevel(logging.DEBUG)

file_handler = logging.FileHandler("logs/debug.log", encoding='utf-8')
file_handler.setLevel(logging.DEBUG)

log_format = logging.Formatter("%(asctime)s - %(name)s - %(levelname)s - %(message)s")
file_handler.setFormatter(log_format)

logger.addHandler(file_handler)

# logger.debug("Debug message")
# logger.info("Info message")
# logger.warning("Warning message")
# logger.error("Error message")
# logger.critical("Critical message")
