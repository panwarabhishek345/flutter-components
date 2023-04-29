<script>
	export let imageURL;
	export let title;
	export let likeCount;
</script>

<div class="col-span-1">
	<div class="bg-white rounded-lg overflow-hidden shadow-md">
        <div class="relative h-0 pb-[80%]">
            <img class="absolute object-cover w-full h-full" src={imageURL} alt={title} />
        </div>
		<div class="px-6 py-4">
			<div class="font-bold text-xl mb-2">{title}</div>
			<p class="text-gray-700 text-base">{likeCount} likes</p>
		</div>
	</div>
</div>
