<script>
	export let imageURL;
</script>

<div class="relative">
	<img src={imageURL} alt="Image 1" class="break-inside-avoid rounded-lg" />
	<div
		class="break-inside-avoid rounded-lg absolute inset-0 flex justify-center items-center bg-black opacity-0 hover:opacity-50 transition duration-300 ease-in-out"
	>
		<button
			class="bg-gray-800 text-white font-bold py-2 px-4 rounded-full mx-2 hover:bg-gray-700 focus:outline-none focus:shadow-outline"
			>Like</button
		>
		<button
			class="bg-gray-800 text-white font-bold py-2 px-4 rounded-full mx-2 hover:bg-gray-700 focus:outline-none focus:shadow-outline"
			>Share</button
		>
	</div>
</div>
