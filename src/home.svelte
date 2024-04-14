<script>
    import { AccordionItem, Accordion } from "flowbite-svelte";
</script>

<head>
    <title> Welcome to Your Fitness Tracker</title>
</head>
<body>
    <h1>Workout Volume Tracker</h1>
    <Accordion>
        <AccordionItem>
            <span slot="header"> About </span>
            <p class="mb-2 text-gray-500 dark:text-gray-400">
                Nicholas and Nathan Workout Tracker
            </p>
        </AccordionItem>
    </Accordion>
</body>

<style>
    body {
        color: #ffffff;
        text-align: center;
        background-color: #333232;
    }
</style>
