import TitleSetter from '@/components/pageTitle';
import { H2 } from '@/components/typography';
import Loader from '@/components/ui/loader';
import { useState } from 'react';
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Card } from '@/components/ui/card';
import { useTables } from '@/context/TableProvider';
import { MejaAddForm } from './AddForm';
import { MejaUpdateForm } from './UpdateForm';
import DeleteMejaDialog from './DeleteDialog';
import { Button } from '@/components/ui/button';
import { SquarePen, Trash } from 'lucide-react';

export default function Page() {
	const { loading, tables } = useTables();
	const [openUpdateDialog, setOpenUpdateDialog] = useState(false);
	const [openDeleteDialog, setOpenDeleteDialog] = useState(false);
	const [updateTarget, setUpdateTarget] = useState<Meja>({} as Meja);

	const triggerUpdateForm = (m: Meja) => {
		setUpdateTarget(m);
		setOpenUpdateDialog(true);
	};

	const triggerDeleteForm = (m: Meja) => {
		setUpdateTarget(m);
		setOpenDeleteDialog(true);
	};

	return loading ? (
		<Loader />
	) : (
		<section className="relative">
			<TitleSetter title="Meja" />
			<div className="-space-y-1 mb-4">
				<H2>Meja</H2>
				<p>Manajemen Meja</p>
			</div>

			<div className="flex justify-end mb-4">
				<MejaAddForm />
				<MejaUpdateForm open={openUpdateDialog} setOpen={setOpenUpdateDialog} updateTarget={updateTarget} />
				<DeleteMejaDialog open={openDeleteDialog} setOpen={setOpenDeleteDialog} id={updateTarget.id} />
			</div>
			<Card className="p-8">
				<Table>
					<TableHeader>
						<TableRow>
							<TableHead>Nomor Meja</TableHead>
							<TableHead>Status</TableHead>
							<TableHead className="w-[10%]"></TableHead>
						</TableRow>
					</TableHeader>
					<TableBody>
						{tables?.map((table) => (
							<TableRow key={table.id}>
								<TableCell>{table.nomor}</TableCell>
								<TableCell>{table.status}</TableCell>
								<TableCell className="w-fit flex gap-4 [&>button]:shadow-md [&>button]:w-12 [&>button]:p-0">
									<Button onClick={() => triggerUpdateForm(table)} variant="outline">
										<SquarePen className="h-4 w-4" />
									</Button>
									<Button onClick={() => triggerDeleteForm(table)} variant="destructive">
										<Trash className="h-4 w-4" />
									</Button>
								</TableCell>
							</TableRow>
						))}
					</TableBody>
				</Table>
			</Card>
		</section>
	);
}
