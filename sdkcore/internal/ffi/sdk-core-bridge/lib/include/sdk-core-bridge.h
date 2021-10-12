#include <stdarg.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>

typedef struct tmprl_activity_heartbeat_t tmprl_activity_heartbeat_t;

typedef struct tmprl_activity_task_completion_t tmprl_activity_task_completion_t;

typedef struct tmprl_activity_task_t tmprl_activity_task_t;

typedef struct tmprl_core_t tmprl_core_t;

typedef struct tmprl_error_t tmprl_error_t;

typedef struct tmprl_log_listener_t tmprl_log_listener_t;

typedef struct tmprl_runtime_t tmprl_runtime_t;

typedef struct tmprl_wf_activation_completion_t tmprl_wf_activation_completion_t;

typedef struct tmprl_wf_activation_t tmprl_wf_activation_t;

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

typedef struct tmprl_bytes {
  const uint8_t *bytes;
  size_t len;
  size_t cap;
} tmprl_bytes;

typedef struct tmprl_worker_config_t {
  /**
   * UTF-8, not null terminated, not owned
   */
  const char *task_queue;
  size_t task_queue_len;
} tmprl_worker_config_t;

/**
 * One and only one field is non-null, caller must free whichever it is
 */
typedef struct tmprl_bytes_or_error_t {
  const struct tmprl_bytes *bytes;
  struct tmprl_error_t *error;
} tmprl_bytes_or_error_t;

/**
 * One and only one field is non-null, caller must free whichever it is
 */
typedef struct tmprl_wf_activation_or_error_t {
  struct tmprl_wf_activation_t *activation;
  struct tmprl_error_t *error;
} tmprl_wf_activation_or_error_t;

/**
 * One and only one field is non-null, caller must free whichever it is
 */
typedef struct tmprl_activity_task_or_error_t {
  struct tmprl_activity_task_t *task;
  struct tmprl_error_t *error;
} tmprl_activity_task_or_error_t;

/**
 * One and only one field is non-null, caller must free error if it's error or
 * pass completion to tmprl_complete_workflow_activation
 */
typedef struct tmprl_wf_activation_completion_or_error_t {
  struct tmprl_wf_activation_completion_t *completion;
  struct tmprl_error_t *error;
} tmprl_wf_activation_completion_or_error_t;

/**
 * One and only one field is non-null, caller must free error if it's error or
 * pass completion to tmprl_complete_activity_task
 */
typedef struct tmprl_activity_task_completion_or_error_t {
  struct tmprl_activity_task_completion_t *completion;
  struct tmprl_error_t *error;
} tmprl_activity_task_completion_or_error_t;

/**
 * One and only one field is non-null, caller must free error if it's error or
 * pass completion to tmprl_record_activity_heartbeat
 */
typedef struct tmprl_activity_heartbeat_or_error_t {
  struct tmprl_activity_heartbeat_t *heartbeat;
  struct tmprl_error_t *error;
} tmprl_activity_heartbeat_or_error_t;

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

void tmprl_bytes_free(struct tmprl_core_t *core, const struct tmprl_bytes *bytes);

struct tmprl_error_t *tmprl_register_worker(struct tmprl_core_t *core,
                                            const struct tmprl_worker_config_t *config);

void tmprl_wf_activation_free(struct tmprl_wf_activation_t *activation);

struct tmprl_bytes_or_error_t tmprl_wf_activation_to_proto(struct tmprl_core_t *core,
                                                           const struct tmprl_wf_activation_t *activation);

/**
 * String is UTF-8, not null terminated, and caller still owns
 */
struct tmprl_wf_activation_or_error_t tmprl_poll_workflow_activation(struct tmprl_core_t *core,
                                                                     const char *task_queue,
                                                                     size_t task_queue_len);

void tmprl_activity_task_free(struct tmprl_activity_task_t *task);

struct tmprl_bytes_or_error_t tmprl_activity_task_to_proto(struct tmprl_core_t *core,
                                                           const struct tmprl_activity_task_t *task);

/**
 * String is UTF-8, not null terminated, and caller still owns
 */
struct tmprl_activity_task_or_error_t tmprl_poll_activity_task(struct tmprl_core_t *core,
                                                               const char *task_queue,
                                                               size_t task_queue_len);

struct tmprl_wf_activation_completion_or_error_t tmprl_wf_activation_completion_from_proto(const uint8_t *bytes,
                                                                                           size_t bytes_len);

/**
 * Takes ownership of the completion and frees internally
 */
struct tmprl_error_t *tmprl_complete_workflow_activation(struct tmprl_core_t *core,
                                                         struct tmprl_wf_activation_completion_t *completion);

struct tmprl_activity_task_completion_or_error_t tmprl_activity_task_completion_from_proto(const uint8_t *bytes,
                                                                                           size_t bytes_len);

/**
 * Takes ownership of the completion and frees internally
 */
struct tmprl_error_t *tmprl_complete_activity_task(struct tmprl_core_t *core,
                                                   struct tmprl_activity_task_completion_t *completion);

struct tmprl_activity_heartbeat_or_error_t tmprl_activity_heartbeat_from_proto(const uint8_t *bytes,
                                                                               size_t bytes_len);

/**
 * Takes ownership of the heartbeat and frees internally
 */
void tmprl_record_activity_heartbeat(struct tmprl_core_t *core,
                                     struct tmprl_activity_heartbeat_t *heartbeat);

/**
 * Strings are UTF-8, not null terminated, and caller still owns
 */
void tmprl_request_workflow_eviction(struct tmprl_core_t *core,
                                     const char *task_queue,
                                     size_t task_queue_len,
                                     const char *run_id,
                                     size_t run_id_len);

void tmprl_shutdown(struct tmprl_core_t *core);

/**
 * String is UTF-8, not null terminated, and caller still owns
 */
void tmprl_shutdown_worker(struct tmprl_core_t *core,
                           const char *task_queue,
                           size_t task_queue_len);

void tmprl_log_listener_free(struct tmprl_log_listener_t *listener);

struct tmprl_log_listener_t *tmprl_log_listener_new(struct tmprl_core_t *core);

/**
 * Blocks and returns true if one was captured or returns false if there will
 * never be another log
 */
bool tmprl_log_listener_next(struct tmprl_core_t *core, struct tmprl_log_listener_t *listener);

const char *tmprl_log_listener_curr_message(const struct tmprl_log_listener_t *listener);

size_t tmprl_log_listener_curr_message_len(const struct tmprl_log_listener_t *listener);

uint64_t tmprl_log_listener_curr_unix_nanos(const struct tmprl_log_listener_t *listener);

size_t tmprl_log_listener_curr_level(const struct tmprl_log_listener_t *listener);
