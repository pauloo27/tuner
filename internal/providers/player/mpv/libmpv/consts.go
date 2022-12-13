package libmpv

//#include <mpv/client.h>
import "C"

//Errors mpv_error
const (
	/**
	 * No error happened (used to signal successful operation).
	 * Keep in mind that many API functions returning error codes can also
	 * return positive values, which also indicate success. API users can
	 * hardcode the fact that ">= 0" means success.
	 */
	ERROR_SUCCESS Error = C.MPV_ERROR_SUCCESS
	/**
	 * The event ringbuffer is full. This means the client is choked, and can't
	 * receive any events. This can happen when too many asynchronous requests
	 * have been made, but not answered. Probably never happens in practice,
	 * unless the mpv core is frozen for some reason, and the client keeps
	 * making asynchronous requests. (Bugs in the client API implementation
	 * could also trigger this, e.g. if events become "lost".)
	 */
	ERROR_EVENT_QUEUE_FULL Error = C.MPV_ERROR_EVENT_QUEUE_FULL
	/**
	 * Memory allocation failed.
	 */
	ERROR_NOMEM Error = C.MPV_ERROR_NOMEM
	/**
	 * The mpv core wasn't configured and initialized yet. See the notes in
	 * mpv_create().
	 */
	ERROR_UNINITIALIZED Error = C.MPV_ERROR_UNINITIALIZED
	/**
	 * Generic catch-all error if a parameter is set to an invalid or
	 * unsupported value. This is used if there is no better error code.
	 */
	ERROR_INVALID_PARAMETER Error = C.MPV_ERROR_INVALID_PARAMETER
	/**
	 * Trying to set an option that doesn't exist.
	 */
	ERROR_OPTION_NOT_FOUND Error = C.MPV_ERROR_OPTION_NOT_FOUND
	/**
	 * Trying to set an option using an unsupported MPV_FORMAT.
	 */
	ERROR_OPTION_FORMAT Error = C.MPV_ERROR_OPTION_FORMAT
	/**
	 * Setting the option failed. Typically this happens if the provided option
	 * value could not be parsed.
	 */
	ERROR_OPTION_ERROR Error = C.MPV_ERROR_OPTION_ERROR
	/**
	 * The accessed property doesn't exist.
	 */
	ERROR_PROPERTY_NOT_FOUND Error = C.MPV_ERROR_PROPERTY_NOT_FOUND
	/**
	 * Trying to set or get a property using an unsupported MPV_FORMAT.
	 */
	ERROR_PROPERTY_FORMAT Error = C.MPV_ERROR_PROPERTY_FORMAT
	/**
	 * The property exists, but is not available. This usually happens when the
	 * associated subsystem is not active, e.g. querying audio parameters while
	 * audio is disabled.
	 */
	ERROR_PROPERTY_UNAVAILABLE Error = C.MPV_ERROR_PROPERTY_UNAVAILABLE
	/**
	 * Error setting or getting a property.
	 */
	ERROR_PROPERTY_ERROR Error = C.MPV_ERROR_PROPERTY_ERROR
	/**
	 * General error when running a command with mpv_command and similar.
	 */
	ERROR_COMMAND Error = C.MPV_ERROR_COMMAND
	/**
	 * Generic error on loading (used with mpv_event_end_file.error).
	 */
	ERROR_LOADING_FAILED Error = C.MPV_ERROR_LOADING_FAILED
	/**
	 * Initializing the audio output failed.
	 */
	ERROR_AO_INIT_FAILED Error = C.MPV_ERROR_AO_INIT_FAILED
	/**
	 * Initializing the video output failed.
	 */
	ERROR_VO_INIT_FAILED Error = C.MPV_ERROR_VO_INIT_FAILED
	/**
	 * There was no audio or video data to play. This also happens if the
	 * file was recognized, but did not contain any audio or video streams,
	 * or no streams were selected.
	 */
	ERROR_NOTHING_TO_PLAY Error = C.MPV_ERROR_NOTHING_TO_PLAY
	/**
	 * When trying to load the file, the file format could not be determined,
	 * or the file was too broken to open it.
	 */
	ERROR_UNKNOWN_FORMAT Error = C.MPV_ERROR_UNKNOWN_FORMAT
	/**
	 * Generic error for signaling that certain system requirements are not
	 * fulfilled.
	 */
	ERROR_UNSUPPORTED Error = C.MPV_ERROR_UNSUPPORTED
	/**
	 * The API function which was called is a stub only.
	 */
	MPV_ERROR_NOT_IMPLEMENTED Error = C.MPV_ERROR_UNSUPPORTED
)

type Format int

//Format mpv_format
const (
	FORMAT_NONE Format = C.MPV_FORMAT_NONE
	/**
	 * The basic type is char*. It returns the raw property string, like
	 * using ${=property} in input.conf (see input.rst).
	 *
	 * NULL isn't an allowed value.
	 *
	 * Warning: although the encoding is usually UTF-8, this is not always the
	 *          case. File tags often store strings in some legacy codepage,
	 *          and even filenames don't necessarily have to be in UTF-8 (at
	 *          least on Linux). If you pass the strings to code that requires
	 *          valid UTF-8, you have to sanitize it in some way.
	 *          On Windows, filenames are always UTF-8, and libmpv converts
	 *          between UTF-8 and UTF-16 when using win32 API functions. See
	 *          the "Encoding of filenames" section for details.
	 *
	 * Example for reading:
	 *
	 *     char *result = NULL;
	 *     if (mpv_get_property(ctx, "property", FORMAT_STRING, = C.MPV_FORMAT_STRING,
	 *         goto error;
	 *     printf("%s\n", result);
	 *     mpv_free(result);
	 *
	 * Or just use mpv_get_property_string().
	 *
	 * Example for writing:
	 *
	 *     char *value = "the new value";
	 *     // yep, you pass the address to the variable
	 *     // (needed for symmetry with other types and mpv_get_property)
	 *     mpv_set_property(ctx, "property", FORMAT_STRING, = C.MPV_FORMAT_STRING,
	 *
	 * Or just use mpv_set_property_string().
	 *
	 */
	FORMAT_STRING Format = C.MPV_FORMAT_STRING
	/**
	 * The basic type is char*. It returns the OSD property string, like
	 * using ${property} in input.conf (see input.rst). In many cases, this
	 * is the same as the raw string, but in other cases it's formatted for
	 * display on OSD. It's intended to be human readable. Do not attempt to
	 * parse these strings.
	 *
	 * Only valid when doing read access. The rest works like MPV_FORMAT_STRING.
	 */
	FORMAT_OSD_STRING Format = C.MPV_FORMAT_OSD_STRING
	/**
	 * The basic type is int. The only allowed values are 0 ("no")
	 * and 1 ("yes").
	 *
	 * Example for reading:
	 *
	 *     int result;
	 *     if (mpv_get_property(ctx, "property", FORMAT_FLAG, = C.MPV_FORMAT_FLAG,
	 *         goto error;
	 *     printf("%s\n", result ? "true" : "false");
	 *
	 * Example for writing:
	 *
	 *     int flag = 1;
	 *     mpv_set_property(ctx, "property", FORMAT_STRING, = C.MPV_FORMAT_STRING,
	 */
	FORMAT_FLAG Format = C.MPV_FORMAT_FLAG
	/**
	 * The basic type is int64_t.
	 */
	FORMAT_INT64 Format = C.MPV_FORMAT_INT64
	/**
	 * The basic type is double.
	 */
	FORMAT_DOUBLE Format = C.MPV_FORMAT_DOUBLE
	/**
	 * The type is mpv_node.
	 *
	 * For reading, you usually would pass a pointer to a stack-allocated
	 * mpv_node value to mpv, and when you're done you call
	 * mpv_free_node_contents(&node).
	 * You're expected not to write to the data - if you have to, copy it
	 * first (which you have to do manually).
	 *
	 * For writing, you construct your own mpv_node, and pass a pointer to the
	 * API. The API will never write to your data (and copy it if needed), so
	 * you're free to use any form of allocation or memory management you like.
	 *
	 * Warning: when reading, always check the mpv_node.format member. For
	 *          example, properties might change their type in future versions
	 *          of mpv, or sometimes even during runtime.
	 *
	 * Example for reading:
	 *
	 *     mpv_node result;
	 *     if (mpv_get_property(ctx, "property", FORMAT_NODE, = C.MPV_FORMAT_NODE,
	 *         goto error;
	 *     printf("format=%d\n", (int)result.format);
	 *     mpv_free_node_contents(&result).
	 *
	 * Example for writing:
	 *
	 *     mpv_node value;
	 *     value.format = MPV_FORMAT_STRING;
	 *     value.u.string = "hello";
	 *     mpv_set_property(ctx, "property", FORMAT_NODE, = C.MPV_FORMAT_NODE,
	 */
	FORMAT_NODE Format = C.MPV_FORMAT_NODE
	/**
	 * Used with mpv_node only. Can usually not be used directly.
	 */
	FORMAT_NODE_ARRAY Format = C.MPV_FORMAT_NODE_ARRAY
	/**
	 * See MPV_FORMAT_NODE_ARRAY.
	 */
	FORMAT_NODE_MAP Format = C.MPV_FORMAT_NODE_MAP
	/**
	 * A raw, untyped byte array. Only used only with mpv_node, and only in
	 * some very special situations. (Currently, only for the screenshot_raw
	 * command.)
	 */
	FORMAT_BYTE_ARRAY = C.MPV_FORMAT_BYTE_ARRAY
)

type EventId int

//EventId  mpv_event_id
const (
	/**
	 * Nothing happened. Happens on timeouts or sporadic wakeups.
	 */
	EVENT_NONE EventId = C.MPV_EVENT_NONE
	/**
	 * Happens when the player quits. The player enters a state where it tries
	 * to disconnect all clients. Most requests to the player will fail, and
	 * mpv_wait_event() will always return instantly (returning new shutdown
	 * events if no other events are queued). The client should react to this
	 * and quit with mpv_detach_destroy() as soon as possible.
	 */
	EVENT_SHUTDOWN EventId = C.MPV_EVENT_SHUTDOWN
	/**
	 * See mpv_request_log_messages().
	 */
	EVENT_LOG_MESSAGE EventId = C.MPV_EVENT_LOG_MESSAGE
	/**
	 * Reply to a mpv_get_property_async() request.
	 * See also mpv_event and mpv_event_property.
	 */
	EVENT_GET_PROPERTY_REPLY EventId = C.MPV_EVENT_GET_PROPERTY_REPLY
	/**
	 * Reply to a mpv_set_property_async() request.
	 * (Unlike EVENT_GET_PROPERTY, = C.MPV_EVENT_GET_PROPERTY,
	 */
	EVENT_SET_PROPERTY_REPLY EventId = C.MPV_EVENT_SET_PROPERTY_REPLY
	/**
	 * Reply to a mpv_command_async() request.
	 */
	EVENT_COMMAND_REPLY EventId = C.MPV_EVENT_COMMAND_REPLY
	/**
	 * Notification before playback start of a file (before the file is loaded).
	 */
	EVENT_START_FILE EventId = C.MPV_EVENT_START_FILE
	/**
	 * Notification after playback end (after the file was unloaded).
	 * See also mpv_event and mpv_event_end_file.
	 */
	EVENT_END_FILE EventId = C.MPV_EVENT_END_FILE
	/**
	 * Notification when the file has been loaded (headers were read etc.), and
	 * decoding starts.
	 */
	EVENT_FILE_LOADED EventId = C.MPV_EVENT_FILE_LOADED
	/**
	 * Idle mode was entered. In this mode, no file is played, and the playback
	 * core waits for new commands. (The command line player normally quits
	 * instead of entering idle mode, unless --idle was specified. If mpv
	 * was started with mpv_create(), idle mode is enabled by default.)
	 */
	EVENT_IDLE EventId = C.MPV_EVENT_IDLE
	/**
	 * Sent every time after a video frame is displayed. Note that currently,
	 * this will be sent in lower frequency if there is no video, or playback
	 * is paused - but that will be removed in the future, and it will be
	 * restricted to video frames only.
	 */
	EVENT_TICK EventId = C.MPV_EVENT_TICK
	/**
	 * Triggered by the script_message input command. The command uses the
	 * first argument of the command as client name (see mpv_client_name()) to
	 * dispatch the message, and passes along all arguments starting from the
	 * second argument as strings.
	 * See also mpv_event and mpv_event_client_message.
	 */
	EVENT_CLIENT_MESSAGE EventId = C.MPV_EVENT_CLIENT_MESSAGE
	/**
	 * Happens after video changed in some way. This can happen on resolution
	 * changes, pixel format changes, or video filter changes. The event is
	 * sent after the video filters and the VO are reconfigured. Applications
	 * embedding a mpv window should listen to this event in order to resize
	 * the window if needed.
	 * Note that this event can happen sporadically, and you should check
	 * yourself whether the video parameters really changed before doing
	 * something expensive.
	 */
	EVENT_VIDEO_RECONFIG EventId = C.MPV_EVENT_VIDEO_RECONFIG
	/**
	 * Similar to EVENT_VIDEO_RECONFIG. = C.MPV_EVENT_VIDEO_RECONFIG.
	 * because there is no such thing as audio output embedding.
	 */
	EVENT_AUDIO_RECONFIG EventId = C.MPV_EVENT_AUDIO_RECONFIG
	/**
	 * Happens when a seek was initiated. Playback stops. Usually it will
	 * resume with EVENT_PLAYBACK_RESTART = C.MPV_EVENT_PLAYBACK_RESTART
	 */
	EVENT_SEEK EventId = C.MPV_EVENT_SEEK
	/**
	 * There was a discontinuity of some sort (like a seek), and playback
	 * was reinitialized. Usually happens after seeking, or ordered chapter
	 * segment switches. The main purpose is allowing the client to detect
	 * when a seek request is finished.
	 */
	EVENT_PLAYBACK_RESTART EventId = C.MPV_EVENT_PLAYBACK_RESTART
	/**
	 * Event sent due to mpv_observe_property().
	 * See also mpv_event and mpv_event_property.
	 */
	EVENT_PROPERTY_CHANGE EventId = C.MPV_EVENT_PROPERTY_CHANGE
	/**
	 * Happens if the internal per-mpv_handle ringbuffer overflows, and at
	 * least 1 event had to be dropped. This can happen if the client doesn't
	 * read the event queue quickly enough with mpv_wait_event(), or if the
	 * client makes a very large number of asynchronous calls at once.
	 *
	 * Event delivery will continue normally once this event was returned
	 * (this forces the client to empty the queue completely).
	 */
	EVENT_QUEUE_OVERFLOW EventId = C.MPV_EVENT_QUEUE_OVERFLOW
	// Internal note: adjust INTERNAL_EVENT_BASE when adding new events.
)

func (eid EventId) String() string {
	switch eid {
	case EVENT_NONE:
		{
			return "EVENT_NONE"
		}
	case EVENT_SHUTDOWN:
		{
			return "EVENT_SHUTDOWN"
		}
	case EVENT_LOG_MESSAGE:
		{
			return "EVENT_LOG_MESSAGE"
		}
	case EVENT_GET_PROPERTY_REPLY:
		{
			return "EVENT_GET_PROPERTY_REPLY"
		}
	case EVENT_SET_PROPERTY_REPLY:
		{
			return "EVENT_SET_PROPERTY_REPLY"
		}
	case EVENT_COMMAND_REPLY:
		{
			return "EVENT_COMMAND_REPLY"
		}
	case EVENT_START_FILE:
		{
			return "EVENT_START_FILE"
		}
	case EVENT_END_FILE:
		{
			return "EVENT_END_FILE"
		}
	case EVENT_FILE_LOADED:
		{
			return "EVENT_FILE_LOADED"
		}
	case EVENT_IDLE:
		{
			return "EVENT_IDLE"
		}
	case EVENT_TICK:
		{
			return "EVENT_TICK"
		}
	case EVENT_CLIENT_MESSAGE:
		{
			return "EVENT_CLIENT_MESSAGE"
		}
	case EVENT_VIDEO_RECONFIG:
		{
			return "EVENT_VIDEO_RECONFIG"
		}
	case EVENT_AUDIO_RECONFIG:
		{
			return "EVENT_AUDIO_RECONFIG"
		}
	case EVENT_SEEK:
		{
			return "EVENT_SEEK"
		}
	case EVENT_PLAYBACK_RESTART:
		{
			return "EVENT_PLAYBACK_RESTART"
		}
	case EVENT_PROPERTY_CHANGE:
		{
			return "EVENT_PROPERTY_CHANGE"
		}
	case EVENT_QUEUE_OVERFLOW:
		{
			return "EVENT_QUEUE_OVERFLOW"
		}
	}
	return "UNKNOWN_EVENT"
}

//Log level  mpv_log_level
const (
	LOG_LEVEL_NONE  = C.MPV_LOG_LEVEL_NONE  /// "no"    - disable absolutely all messages
	LOG_LEVEL_FATAL = C.MPV_LOG_LEVEL_FATAL /// "fatal" - critical/aborting errors
	LOG_LEVEL_ERROR = C.MPV_LOG_LEVEL_ERROR /// "error" - simple errors
	LOG_LEVEL_WARN  = C.MPV_LOG_LEVEL_WARN  /// "warn"  - possible problems
	LOG_LEVEL_INFO  = C.MPV_LOG_LEVEL_INFO  /// "info"  - informational message
	LOG_LEVEL_V     = C.MPV_LOG_LEVEL_V     /// "v"     - noisy informational message
	LOG_LEVEL_DEBUG = C.MPV_LOG_LEVEL_DEBUG /// "debug" - very noisy technical information
	LOG_LEVEL_TRACE = C.MPV_LOG_LEVEL_TRACE /// "trace" - extremely noisy
)

type EndFileReason int

//EndFileReason mpv_end_file_reason
const (
	/**
	 * The end of file was reached. Sometimes this may also happen on
	 * incomplete or corrupted files, or if the network connection was
	 * interrupted when playing a remote file. It also happens if the
	 * playback range was restricted with --end or --frames or similar.
	 */
	END_FILE_REASON_EOF EndFileReason = C.MPV_END_FILE_REASON_EOF
	/**
	 * Playback was stopped by an external action (e.g. playlist controls).
	 */
	END_FILE_REASON_STOP EndFileReason = C.MPV_END_FILE_REASON_STOP
	/**
	 * Playback was stopped by the quit command or player shutdown.
	 */
	END_FILE_REASON_QUIT EndFileReason = C.MPV_END_FILE_REASON_QUIT
	/**
	 * Some kind of error happened that lead to playback abort. Does not
	 * necessarily happen on incomplete or broken files (in these cases, both
	 * MPV_END_FILE_REASON_ERROR or MPV_END_FILE_REASON_EOF are possible).
	 *
	 * mpv_event_end_file.error will be set.
	 */
	END_FILE_REASON_ERROR EndFileReason = C.MPV_END_FILE_REASON_ERROR
	/**
	 * The file was a playlist or similar. When the playlist is read, its
	 * entries will be appended to the playlist after the entry of the current
	 * file, the entry of the current file is removed, and a MPV_EVENT_END_FILE
	 * event is sent with reason set to MPV_END_FILE_REASON_REDIRECT. Then
	 * playback continues with the playlist contents.
	 * Since API version 1.18.
	 */
	END_FILE_REASON_REDIRECT EndFileReason = C.MPV_END_FILE_REASON_REDIRECT
)

func (efr EndFileReason) String() string {
	switch efr {
	case END_FILE_REASON_EOF:
		return "END_FILE_REASON_EOF"
	case END_FILE_REASON_STOP:
		return "END_FILE_REASON_STOP"
	case END_FILE_REASON_QUIT:
		return "END_FILE_REASON_QUIT"
	case END_FILE_REASON_ERROR:
		return "END_FILE_REASON_ERROR"
	case END_FILE_REASON_REDIRECT:
		return "END_FILE_REASON_REDIRECT"
	default:
		return "END_FILE_REASON_UNKNOWN"
	}
}
