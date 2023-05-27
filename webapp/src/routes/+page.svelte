<script lang="ts">
    const exercices = [
        "verbes"
    ];
</script>

<h1>Click on a course to do the exercices</h1>
<ul>
    {#each exercices as exercice}
        <li><a href="/exercice/?name={exercice}">{exercice}</a></li>
    {/each}
</ul>