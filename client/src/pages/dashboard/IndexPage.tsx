import TitleSetter from '@/components/pageTitle';
import { Card } from '@/components/ui/card';

export default function IndexPage() {
  return (
    <>
      <TitleSetter title="Dashboard" />
      <div className="grid grid-cols-2 lg:grid-cols-4 gap-4 lg:gap-8">
        {Array.from({ length: 4 }, (_, index) => (
          <OverviewCard key={index} />
        ))}
      </div>
    </>
  );
}

const OverviewCard = () => {
  return (
    <Card className="h-[170px] flex place-items-center">
      <p className="mx-auto font-medium text-primary">This is card</p>
    </Card>
  );
};
