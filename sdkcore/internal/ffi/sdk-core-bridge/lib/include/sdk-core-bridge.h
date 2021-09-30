#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct tmprl_core_t tmprl_core_t;

typedef struct tmprl_error_t tmprl_error_t;

typedef struct tmprl_runtime_t tmprl_runtime_t;

/**
 * One and only one field is non-null, caller must free whichever it is
 */
typedef struct tmprl_core_or_error_t {
  struct tmprl_core_t *core;
  struct tmprl_error_t *error;
} tmprl_core_or_error_t;

typedef struct tmprl_core_new_options_t {
  /**
   * UTF-8, not null terminated, not owned
   */
  const char *target_url;
  size_t target_url_len;
} tmprl_core_new_options_t;

typedef struct tmprl_core_worker_config_t {
  /**
   * UTF-8, not null terminated, not owned
   */
  const char *task_queue;
  size_t task_queue_len;
} tmprl_core_worker_config_t;

void hello_rust(void);

void tmprl_error_free(struct tmprl_error_t *error);

/**
 * Not null terminated (use tmprl_error_message_len()) and remains owned by
 * error
 */
const char *tmprl_error_message(const struct tmprl_error_t *error);

size_t tmprl_error_message_len(const struct tmprl_error_t *error);

int tmprl_error_message_code(const struct tmprl_error_t *error);

struct tmprl_runtime_t *tmprl_runtime_new(void);

void tmprl_runtime_free(struct tmprl_runtime_t *runtime);

struct tmprl_core_or_error_t tmprl_core_new(struct tmprl_runtime_t *runtime,
                                            const struct tmprl_core_new_options_t *options);

void tmprl_core_free(struct tmprl_core_t *core);

struct tmprl_error_t *tmprl_core_register_worker(struct tmprl_core_t *core,
                                                 const struct tmprl_core_worker_config_t *config);
