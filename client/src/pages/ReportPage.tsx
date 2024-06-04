import TitleSetter from '@/components/pageTitle';
import { Button } from '@/components/ui/button';
import { Loader2 } from 'lucide-react';
import { useState } from 'react';

export default function ReportPage() {
  const [loading, setLoading] = useState(false);
  const downloadExcel = async () => {
    try {
      setLoading(true);
      const response = await fetch('/api/v1/reports', {
        method: 'GET',
      });

      if (!response.ok) {
        throw new Error('Network response was not ok');
      }

      const blob = await response.blob();

      // Create a link element
      const link = document.createElement('a');

      // Set the download attribute with a filename
      link.href = window.URL.createObjectURL(blob);
      link.download = 'report.xlsx';

      // Append the link to the body
      document.body.appendChild(link);

      // Programmatically click the link to trigger the download
      link.click();

      // Remove the link from the document
      document.body.removeChild(link);
    } catch (error) {
      console.error('Error downloading the Excel file:', error);
    } finally {
      setLoading(false);
    }
  };
  return (
    <>
      <TitleSetter title="Laporan" />
      <p>This Is Laporan Page</p>
    </>
  );
}
// 			<Button disabled={loading} onClick={downloadExcel}>
// 				{loading ? <Loader2 className="animate-spin" /> : <span>Download</span>}
// 			</Button>
