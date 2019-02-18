CONCURSE WORKER HEALTH CHECKER

	A "prober" that tries to create volumes and containers to
	assess whether a Concourse worker is healthy or not.


HOW IT WORKS


	--> check_health
		--> baggageclaim
			--> create_volume(handle)
			<-- OK
			--> garden
				--> create_container(handle, volume_location)
				<-- OK
			<-- OK

