<script lang="ts">
	import { type Snippet, onMount } from 'svelte'

	import { IsLaunched } from '@wailsjs/go/app/App.js'
	import { EventsOn } from '@wailsjs/runtime/runtime'

	import { store } from './MainStore.svelte'

	interface Props {
		children?: Snippet
	}
	let { children }: Props = $props()
	const id = $props.id()
	store.drawerId = `main-drawer-${id}`

	EventsOn('launched', () => (store.launched = true))
	onMount(async () => (store.launched = await IsLaunched()))
</script>

<div class="drawer-open drawer size-full">
	<input id={store.drawerId} type="checkbox" class="drawer-toggle" />
	<div class="drawer-content overflow-hidden">
		{@render children?.()}
	</div>
	<div class="drawer-side is-drawer-close:overflow-visible">
		<aside
			class="flex h-full flex-col items-start bg-base-200 shadow-inner is-drawer-close:w-14 is-drawer-open:w-64"
		>
			<ul class="menu w-full grow">
				<li>
					<a
						class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
						data-tip="メイン"
						href="/"
					>
						<i class="i-material-symbols-map"></i>
						<span class="is-drawer-close:hidden">メイン</span>
					</a>
				</li>
				<li>
					<button
						class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
						data-tip="インポート"
					>
						<i class="i-material-symbols-file-open"></i>
						<span class="is-drawer-close:hidden">インポート</span>
					</button>
				</li>
				<li>
					<button
						class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
						data-tip="レイヤー削除"
					>
						<i class="i-material-symbols-delete"></i>
						<span class="is-drawer-close:hidden">レイヤー削除</span>
					</button>
				</li>
				<li>
					<a
						class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
						data-tip="住所リスト"
						href="/address"
					>
						<i class="i-material-symbols-database"></i>
						<span class="is-drawer-close:hidden">住所リスト</span>
					</a>
				</li>
				<li>
					<a
						class="is-drawer-close:tooltip is-drawer-close:tooltip-right"
						data-tip="年賀状"
						href="/nenga"
					>
						<i class="i-material-symbols-mail"></i>
						<span class="is-drawer-close:hidden">年賀状</span>
					</a>
				</li>
			</ul>
			<div class="flex p-2">
				<label for={store.drawerId} class="btn btn-circle btn-ghost">
					<i class="i-material-symbols-menu"></i>
				</label>
			</div>
		</aside>
	</div>
</div>
