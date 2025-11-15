package capture

func JSEnsure() string {
	return `

		if (waitForQueueAsync===undefined) {
			function waitForQueueAsync() {
				return Promise.resolve();
			}
		}

		(async () => {
			console.log("Start");

			await waitForQueueAsync();

			console.log("Runs after stack is clear");
		})();

		"OK"
	`
}

func JSSetFrame(f string) string {

	return `
		window.context.setFrame(`+f+`);
		"OK";
	` 

}