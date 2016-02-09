java -jar scheduler-0.7.0.jar --zookeeperMesosUrl zk://172.31.44.94:2181,172.31.38.145:2181,172.31.40.63:2181/mesos --elasticsearchCpu 0.5 --executorForcePullImage true


In ElasticsearchScheduler.java (https://github.com/dvonthenen/elasticsearch/blob/master/scheduler/src/main/java/org/apache/mesos/elasticsearch/scheduler/ElasticsearchScheduler.java)
  public void resourceOffers(SchedulerDriver driver, List<Protos.Offer> offers)
    Call taskInfoFactory.createTask()
      In TaskInfoFactory.java (https://github.com/dvonthenen/elasticsearch/blob/master/scheduler/src/main/java/org/apache/mesos/elasticsearch/scheduler/TaskInfoFactory.java)
        CommandInfo???
OR
        ExecutorInfo??? (This looks more correct)


https://github.com/dvonthenen/elasticsearch/blob/master/scheduler/src/main/java/org/apache/mesos/elasticsearch/scheduler/TaskInfoFactory.java
private Protos.CommandInfo.Builder newCommandInfo(Configuration configuration)

To Add Environment Variables




https://github.com/apache/mesos/blob/master/include/mesos/mesos.proto

  Describes a task. Passed from the scheduler all the way to an
  executor (see SchedulerDriver::launchTasks and
  Executor::launchTask). Either ExecutorInfo or CommandInfo should be set.
  A different executor can be used to launch this task, and subsequent tasks
  meant for the same executor can reuse the same ExecutorInfo struct.

message TaskInfo {
  required string name = 1;
  required TaskID task_id = 2;
  required SlaveID slave_id = 3;
  repeated Resource resources = 4;
  optional ExecutorInfo executor = 5;
  optional CommandInfo command = 7;
  // Task provided with a container will launch the container as part
  // of this task paired with the task's CommandInfo.
  optional ContainerInfo container = 9;
  optional bytes data = 6;
  // A health check for the task (currently in *alpha* and initial
  // support will only be for TaskInfo's that have a CommandInfo).
  optional HealthCheck health_check = 8;

  // Labels are free-form key value pairs which are exposed through
  // master and slave endpoints. Labels will not be interpreted or
  // acted upon by Mesos itself. As opposed to the data field, labels
  // will be kept in memory on master and slave processes. Therefore,
  // labels should be used to tag tasks with light-weight meta-data.
  optional Labels labels = 10;

  // Service discovery information for the task. It is not interpreted
  // or acted upon by Mesos. It is up to a service discovery system
  // to use this information as needed and to handle tasks without
  // service discovery information.
  optional DiscoveryInfo discovery = 11;
}


  Describes a command, executed via: '/bin/sh -c value'. Any URIs specified
  are fetched before executing the command.  If the executable field for an
  uri is set, executable file permission is set on the downloaded file.
  Otherwise, if the downloaded file has a recognized archive extension
  (currently [compressed] tar and zip) it is extracted into the executor's
  working directory. This extraction can be disabled by setting `extract` to
  false. In addition, any environment variables are set before executing
  the command (so they can be used to "parameterize" your command).

message CommandInfo {
  message URI {
    required string value = 1;
    optional bool executable = 2;

    // In case the fetched file is recognized as an archive, extract
    // its contents into the sandbox. Note that a cached archive is
    // not copied from the cache to the sandbox in case extraction
    // originates from an archive in the cache.
    optional bool extract = 3 [default = true];

    // If this field is "true", the fetcher cache will be used. If not,
    // fetching bypasses the cache and downloads directly into the
    // sandbox directory, no matter whether a suitable cache file is
    // available or not. The former directs the fetcher to download to
    // the file cache, then copy from there to the sandbox. Subsequent
    // fetch attempts with the same URI will omit downloading and copy
    // from the cache as long as the file is resident there. Cache files
    // may get evicted at any time, which then leads to renewed
    // downloading. See also "docs/fetcher.md" and
    // "docs/fetcher-cache-internals.md".
    optional bool cache = 4;
  }

  repeated URI uris = 1;

  optional Environment environment = 2;

  // There are two ways to specify the command:
  // 1) If 'shell == true', the command will be launched via shell
  //		(i.e., /bin/sh -c 'value'). The 'value' specified will be
  //		treated as the shell command. The 'arguments' will be ignored.
  // 2) If 'shell == false', the command will be launched by passing
  //		arguments to an executable. The 'value' specified will be
  //		treated as the filename of the executable. The 'arguments'
  //		will be treated as the arguments to the executable. This is
  //		similar to how POSIX exec families launch processes (i.e.,
  //		execlp(value, arguments(0), arguments(1), ...)).
  // NOTE: The field 'value' is changed from 'required' to 'optional'
  // in 0.20.0. It will only cause issues if a new framework is
  // connecting to an old master.
  optional bool shell = 6 [default = true];
  optional string value = 3;
  repeated string arguments = 7;

  // Enables executor and tasks to run as a specific user. If the user
  // field is present both in FrameworkInfo and here, the CommandInfo
  // user value takes precedence.
  optional string user = 5;
}

message Environment {
  message Variable {
    required string name = 1;
    required string value = 2;
  }

  repeated Variable variables = 1;
}
