// import ExpensesHeader from "../components/navigation/ExpensesHeader";
import Statistics from "../components/system/Statistics";
import Chart from "../components/system/Chart";


const DUMMY_EXPENSES = [
    {
        id: 1,
        title: "First",
        amount: 12.99,
        date: new Date().toISOString()
    },
    {
        id: 2,
        title: "Second",
        amount: 5.67,
        date: new Date().toISOString()
    },
    {
        id: 3,
        title: "Third",
        amount: 23.99,
        date: new Date().toISOString()
    }
];

export default function AnalysisPage() {
    return(
        <>
            {/* <ExpensesHeader /> */}
            <main>
                <Chart expenses={ DUMMY_EXPENSES } />
                <Statistics expenses={ DUMMY_EXPENSES } />
            </main>
        </>
        
    )
}
