<script lang="ts">
	import { onMount } from 'svelte'
	import { AllNengaAtena, SaveNengaAtena, DeleteNengaAtena } from '@wailsjs/go/app/App'
	import type { domain } from '@wailsjs/go/models'

	interface NengaEntry {
		postcode: string
		address: string
		names: string[]
	}
	interface NengaData {
		sender: NengaEntry | null
		destinations: NengaEntry[]
	}
	interface NengaAtenaRaw {
		ID: string
		Year: number
		Names: string[]
		Postcode: string
		Address: string
		GroupName: string
		Kind: number
		IsSender: boolean
	}

	let year = $state(2026)
	let nengaData = $state<domain.NengaAtena | null>(null)
	let rawList = $state<NengaAtenaRaw[]>([])
	let loading = $state(false)
	let activeTab = $state<'preview' | 'edit'>('preview')
	let error = $state('')

	// 新規登録用
	let newNames = $state('')
	let newPostcode = $state('')
	let newAddress = $state('')
	let newGroupName = $state('')
	let newIsSender = $state(false)

	async function loadData() {
		loading = true
		error = ''
		try {
			nengaData = await AllNengaAtena(year)
		} catch (e: any) {
			error = e?.message || String(e)
		} finally {
			loading = false
		}
	}

	async function onAdd() {
		const nameArray = newNames.split('\n').map((n) => n.trim()).filter((n) => n)
		if (nameArray.length === 0) return
		try {
			await SaveNengaAtena({year, Names: nameArray, newPostcode, newAddress, newGroupName, newIsSender})
			newNames = ''
			newPostcode = ''
			newAddress = ''
			newGroupName = ''
			newIsSender = false
			await loadData()
		} catch (e: any) {
			error = e?.message || String(e)
		}
	}

	async function onDelete(id: string) {
		try {
			await DeleteNengaAtena(id)
			await loadData()
		} catch (e: any) {
			error = e?.message || String(e)
		}
	}

	function printPreview() {
		window.print()
	}

	onMount(async () => {
		year = 2026
		await loadData()
	})
</script>

<svelte:head>
	<style>
		@media print {
			/* 印刷時はプレビュー以外を隠す */
			body * {
				visibility: hidden;
			}
			.print-area,
			.print-area * {
				visibility: visible;
			}
			.print-area {
				position: absolute;
				left: 0;
				top: 0;
			}
		}
	</style>
</svelte:head>

<div class="flex size-full flex-col overflow-hidden">
	<!-- ヘッダー : 印刷時は非表示 -->
	<div class="flex items-center gap-2 border-b border-base-300 p-3 print:hidden">
		<h1 class="text-lg font-bold">年賀状 宛名面</h1>
		<div class="flex items-center gap-1">
			<input
				type="number"
				class="input input-sm w-24"
				bind:value={year}
				min="2000"
				max="2100"
			/>
			<span class="text-sm">年</span>
		</div>
		<button class="btn btn-primary btn-sm" onclick={() => loadData()} disabled={loading}>
			{#if loading}
				<span class="loading loading-spinner loading-xs"></span>
			{/if}
			読込
		</button>
		<div class="grow"></div>
		<div role="tablist" class="tabs tabs-box tabs-sm">
			<button
				role="tab"
				class="tab"
				class:tab-active={activeTab === 'preview'}
				onclick={() => (activeTab = 'preview')}
			>
				プレビュー ({nengaData?.destinations?.length ?? 0}件)
			</button>
			<button
				role="tab"
				class="tab"
				class:tab-active={activeTab === 'edit'}
				onclick={() => (activeTab = 'edit')}
			>
				編集 ({rawList.length}件)
			</button>
		</div>
		<button class="btn btn-secondary btn-sm" onclick={printPreview}>
			<i class="i-material-symbols-print"></i>
			印刷
		</button>
	</div>

	{#if error}
		<div class="alert alert-error m-2 print:hidden">
			<span>{error}</span>
			<button class="btn btn-ghost btn-xs" onclick={() => (error = '')}>✕</button>
		</div>
	{/if}

	<!-- プレビュータブ -->
	{#if activeTab === 'preview'}
		<div class="print-area flex-1 overflow-y-auto bg-base-200 p-4 print:overflow-visible print:bg-transparent print:p-0">
			{#if nengaData}
				{#each nengaData.destinations as dest}
					<section class="sheet">
						<p class="郵便番号">{dest.postcode}</p>
						<p class="住所">{@html dest.address}</p>
						<div class="宛名">
							{#each dest.names as name}
								<p>{name}　様</p>
							{/each}
						</div>
						{#if nengaData.sender}
							<div class="差出人住所氏名">
								<p class="差出人住所">{@html nengaData.sender.address}</p>
								<p class="差出人">
									{#each nengaData.sender.names as name}
										{name}<br />
									{/each}
								</p>
							</div>
							<p class="差出人郵便番号">{nengaData.sender.postcode}</p>
						{/if}
					</section>
				{/each}
				{#if nengaData.destinations.length === 0}
					<div class="flex flex-col items-center justify-center gap-4 py-16 text-base-content/50">
						<i class="i-material-symbols-mail text-6xl"></i>
						<p class="text-lg">宛名データがありません</p>
						<p class="text-sm">「編集」タブから宛名を追加してください</p>
					</div>
				{/if}
			{/if}
		</div>
	{/if}

	<!-- 編集タブ -->
	{#if activeTab === 'edit'}
		<div class="flex flex-1 flex-col gap-4 overflow-y-auto p-4">
			<!-- 新規追加フォーム -->
			<div class="card bg-base-100 shadow">
				<div class="card-body gap-3 p-4">
					<h2 class="card-title text-sm">新規追加</h2>
					<div class="grid grid-cols-2 gap-3">
						<div class="form-control">
							<label class="label" for="new-postcode">
								<span class="label-text">郵便番号</span>
							</label>
							<input
								id="new-postcode"
								type="text"
								class="input input-bordered input-sm"
								placeholder="1234567"
								bind:value={newPostcode}
							/>
						</div>
						<div class="form-control">
							<label class="label" for="new-group">
								<span class="label-text">グループ名</span>
							</label>
							<input
								id="new-group"
								type="text"
								class="input input-bordered input-sm"
								bind:value={newGroupName}
							/>
						</div>
					</div>
					<div class="form-control">
						<label class="label" for="new-address">
							<span class="label-text">住所</span>
						</label>
						<input
							id="new-address"
							type="text"
							class="input input-bordered input-sm"
							placeholder="東京都千代田区1-1-1"
							bind:value={newAddress}
						/>
					</div>
					<div class="form-control">
						<label class="label" for="new-names">
							<span class="label-text">名前（改行区切り）</span>
						</label>
						<textarea
							id="new-names"
							class="textarea textarea-bordered textarea-sm"
							rows="2"
							placeholder="山田 太郎&#10;山田 花子"
							bind:value={newNames}
						></textarea>
					</div>
					<div class="flex items-center gap-4">
						<label class="label cursor-pointer gap-2">
							<input type="checkbox" class="checkbox checkbox-sm" bind:checked={newIsSender} />
							<span class="label-text">差出人として登録</span>
						</label>
						<div class="grow"></div>
						<button class="btn btn-primary btn-sm" onclick={onAdd}>追加</button>
					</div>
				</div>
			</div>

			<!-- 一覧 -->
			<div class="overflow-x-auto">
				<table class="table table-zebra table-sm">
					<thead>
						<tr>
							<th>区分</th>
							<th>〒</th>
							<th>住所</th>
							<th>名前</th>
							<th>グループ</th>
							<th></th>
						</tr>
					</thead>
					<tbody>
						{#each rawList as item}
							<tr>
								<td>
									{#if item.IsSender}
										<span class="badge badge-secondary badge-sm">差出人</span>
									{:else}
										<span class="badge badge-primary badge-sm">宛先</span>
									{/if}
								</td>
								<td class="font-mono text-xs">{item.Postcode}</td>
								<td class="max-w-48 truncate text-xs">{item.Address}</td>
								<td class="text-xs">{item.Names?.join(', ')}</td>
								<td class="text-xs">{item.GroupName}</td>
								<td>
									<button
										class="btn btn-error btn-ghost btn-xs"
										onclick={() => onDelete(item.ID)}
										aria-label="削除"
									>
										<i class="i-material-symbols-delete"></i>
									</button>
								</td>
							</tr>
						{/each}
					</tbody>
				</table>
			</div>
		</div>
	{/if}
</div>

<style>
	/* 年賀はがき宛名面のスタイル（nenga.css / nenga.tpl から移植） */
	@page {
		size: 100mm 148mm;
		margin: 0;
		padding: 0;
	}

	.sheet {
		width: 100mm;
		height: 148mm;
		margin: 0 auto 16px;
		padding: 0;
		position: relative;
		page-break-after: always;
		font-family: 'HiraMinProN-W3', 'IPAmj明朝', serif;
		font-style: normal;
		font-weight: 400;
	}

	.sheet :global(p) {
		margin: 0;
		padding: 0;
	}

	:global(.郵便番号) {
		font-size: 12pt;
		margin: 0;
		padding: 0;
		position: absolute;
		top: 12mm;
		left: 46mm;
		letter-spacing: 4.8mm;
	}

	:global(.住所) {
		font-size: 14pt;
		margin-left: 0px;
		padding-top: 0px;
		padding-bottom: 20mm;
		line-height: 1.5em;
		writing-mode: vertical-rl;
		-webkit-writing-mode: vertical-rl;
		text-orientation: upright;
		-webkit-text-orientation: upright;
		position: absolute;
		top: 30mm;
		right: 8mm;
	}

	:global(.宛名) {
		font-size: 22pt;
		letter-spacing: 0.1em;
		position: absolute;
		top: 30%;
		left: 57%;
		-webkit-transform: translateY(-10%) translateX(-50%);
		transform: translateY(-10%) translateX(-50%);
		writing-mode: vertical-rl;
		-webkit-writing-mode: vertical-rl;
		text-orientation: upright;
		-webkit-text-orientation: upright;
	}

	:global(.差出人住所氏名) {
		position: absolute;
		top: 70mm;
		left: 3mm;
		writing-mode: vertical-rl;
		-webkit-writing-mode: vertical-rl;
		text-orientation: upright;
		-webkit-text-orientation: upright;
		line-height: 1.2em;
	}

	:global(.差出人住所) {
		font-size: 10pt;
		margin: 0;
	}

	:global(.差出人) {
		font-size: 12pt;
		letter-spacing: 0.1em;
		margin: 0;
		padding-top: 10mm;
	}

	:global(.差出人郵便番号) {
		font-size: 10pt;
		margin: 0;
		padding: 0;
		position: absolute;
		top: 124mm;
		left: 6mm;
		letter-spacing: 2.1mm;
	}

	/* 画面表示時のスタイル */
	@media screen {
		.sheet {
			background: white;
			box-shadow: 0 0.5mm 2mm rgba(0, 0, 0, 0.3);
		}
	}

	@media print {
		.sheet {
			margin: 0;
		}
	}
</style>
